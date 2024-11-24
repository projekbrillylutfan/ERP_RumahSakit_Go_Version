package app

import "github.com/go-redis/redis/v8"

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Addr:     "localhost:6379", // ganti sesuai kebutuhan
		Password: "",              // kosong jika tidak ada password
		DB:       0,               // gunakan database redis default
	}
}

func NewRedisClient(cfg *RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}