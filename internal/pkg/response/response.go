package response

import (
	"net/http"

	"github.com/bytedance/sonic"
)

type Meta struct {
	Total int `json:"total,omitempty"`
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"`
}

func OK(m string, data any, meta *Meta) *Response {
	return &Response{
		Status:  "OK",
		Message: m,
		Data:    data,
		Meta:    meta,
	}
}

func Failed(m string, err error) *Response {
	return &Response{
		Status:  "FAILED",
		Message: m,
		Error:   err.Error(),
	}
}

func Created(m string, data any) *Response {
	return &Response{
		Status:  "CREATED",
		Message: m,
		Data:    data,
	}
}

func (r *Response) JSON(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp, err := sonic.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Response) Text(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)

	text := r.Message
	if r.Error != "" {
		text = r.Error
	}

	_, err := w.Write([]byte(text))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
