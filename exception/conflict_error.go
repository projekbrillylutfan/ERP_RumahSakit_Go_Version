package exception

type ConflictError struct {
	Message string
}

func (conflictError ConflictError) Error() string {
	return conflictError.Message
}