package errors

type Response struct {
	Status  int
	Message string
	Error   error
}
