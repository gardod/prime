package response

import (
	"net/http"
)

func Prepare(w http.ResponseWriter, v interface{}, code int) *Response {
	resp := NewResponse(w)
	resp.SetStatusCode(code)

	if err, ok := v.(error); ok {
		resp.Error = err.Error()
	} else {
		resp.Data = v
	}

	return resp
}

func JSON(w http.ResponseWriter, v interface{}, code int) {
	Prepare(w, v, code).JSON()
}
