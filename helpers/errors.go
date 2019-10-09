package helpers

type AdaptorError struct {
	Code    int
	Message string
}

func (err AdaptorError) Error() string {
	return err.Message
}

func (err AdaptorError) StatusCode() int {
	return err.Code
}
