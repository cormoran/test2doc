package test

import (
	"net/http"
	"net/http/httptest"
)

type ResponseWriter struct {
	HandlerInfo HandlerInfo
	URLVars     map[string]string
	W           *httptest.ResponseRecorder
}

type HandlerInfo struct {
	FileName string
	FuncName string
}

func NewResponseWriter(w *httptest.ResponseRecorder) *ResponseWriter {
	return &ResponseWriter{
		W: w,
	}
}

func (rw *ResponseWriter) Header() http.Header {
	return rw.W.Header()
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	rw.setHandlerInfo()
	return rw.W.Write(b)
}

func (rw *ResponseWriter) WriteHeader(c int) {
	rw.W.WriteHeader(c)
}

func (rw *ResponseWriter) setHandlerInfo() {
	rw.HandlerInfo = GetHandlerInfo()
}
