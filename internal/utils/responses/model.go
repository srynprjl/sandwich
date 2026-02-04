package responses

type Response struct {
	Status  int
	Message string
	Error   error
}

func (r Response) WebResponse() map[string]any {
	return map[string]any{"status": r.Status, "message": r.Message}
}
