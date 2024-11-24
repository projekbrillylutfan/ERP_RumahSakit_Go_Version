package exception

import "strconv"

func ConversionErrorStrconv(id string) int64 {
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(ConversionError{
			Message: err.Error(),
		})
	}

	return idInt64
}