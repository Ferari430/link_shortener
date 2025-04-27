package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	Code int
}

func (w *WrapperWriter) Wrapper(statuscode int) {
	w.ResponseWriter.WriteHeader(statuscode)
	w.Code = statuscode
}
