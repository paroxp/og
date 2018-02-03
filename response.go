package main

type Response struct {
	Type    string            `json:"type,omitempty"`
	Action  string            `json:"action,omitempty"`
	Body    interface{}       `json:"body,omitempty"`
	Message string            `json:"message,omitempty"`
	Meta    map[string]string `json:"meta,omitempty"`
}

func NewErrorResponse(err error) *Response {
	return &Response{
		Type:    "error",
		Message: err.Error(),
	}
}
