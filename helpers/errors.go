package helpers

type DBError struct {
	Code    int
	Message string
}

func (err DBError) Error() string {
	return err.Message
}

func (err DBError) StatusCode() int {
	return err.Code
}
