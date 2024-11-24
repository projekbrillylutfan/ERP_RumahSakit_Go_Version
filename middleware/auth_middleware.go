package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

func AuthenticateJWT(allowedRoles []string, config configuration.Config) func(*fiber.Ctx) error {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return func(ctx *fiber.Ctx) error {
		// Ambil token dari Authorization header
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return ctx.
				Status(fiber.StatusBadRequest).
				JSON(&dto.GeneralResponseError{
					Code:    400,
					Message: "Bad Request",
					Error:    "Missing Authorization header",
				})
		}

		// Pastikan format header adalah "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return ctx.
				Status(fiber.StatusBadRequest).
				JSON(&dto.GeneralResponseError{
					Code:    400,
					Message: "Bad Request",
					Error:    "Invalid Authorization header format",
				})
		}

		tokenString := parts[1]

		// Parse dan validasi token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Validasi algoritma
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(&dto.GeneralResponseError{
					Code:    401,
					Message: "Unauthorized",
					Error:    "Invalid or expired JWT",
				})
		}

		// Ambil claims dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(&dto.GeneralResponseError{
					Code:    401,
					Message: "Unauthorized",
					Error:    "Invalid JWT claims",
				})
		}

		// Ambil roles dari claims
		roles, ok := claims["roles"].(string)
		if !ok {
			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(&dto.GeneralResponseError{
					Code:    401,
					Message: "Unauthorized",
					Error:    "Invalid roles in JWT",
				})
		}

		// Log role untuk debugging
		configuration.NewLogger().Info("Allowed roles: ", allowedRoles, " User role: ", roles)

		// Cek apakah role pengguna ada di daftar role yang diizinkan
		isAllowed := false
		for _, allowedRole := range allowedRoles {
			if roles == allowedRole {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			return ctx.Next() // Jika cocok, lanjutkan ke handler berikutnya
		}

		return ctx.
			Status(fiber.StatusUnauthorized).
			JSON(&dto.GeneralResponseError{
				Code:    401,
				Message: "Unauthorized",
				Error:    "Invalid Role",
			})
	}
}
