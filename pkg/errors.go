package pkg

type ErrorPayload struct {
	Error string `json:"error"`
}

func ErrorResponse(err error) ErrorPayload {
	return ErrorPayload{
		Error: err.Error(),
	}
}
