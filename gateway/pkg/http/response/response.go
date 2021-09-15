package response

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Response struct {
	w http.ResponseWriter

	cookies []*http.Cookie
	code    int

	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w: w}
}

func (r *Response) AddCookie(c *http.Cookie) {
	r.cookies = append(r.cookies, c)
}

func (r *Response) SetStatusCode(code int) {
	r.code = code
	r.Status = http.StatusText(r.code)
}

func (r *Response) JSON() {
	r.writeHeader("application/json")

	err := json.NewEncoder(r.w).Encode(r)
	if err != nil {
		logrus.WithError(err).Error("Unable to encode Response")
	}
}

func (r *Response) writeHeader(contentType string) {
	r.w.Header().Add("Content-Type", contentType)
	for _, c := range r.cookies {
		http.SetCookie(r.w, c)
	}
	r.w.WriteHeader(r.code)
}
