package exception

type ConversionError struct {
	Message string
}

func (conversionError ConversionError) Error() string {
	return conversionError.Message
}