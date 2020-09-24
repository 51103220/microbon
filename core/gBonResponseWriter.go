package core

import "net/http"

type GBonResponseWriter struct {
	HeaderWritten bool
	Done          bool
	http.ResponseWriter
	StatusCode int
}

func (w *GBonResponseWriter) WriteHeader(statusCode int) {
	w.HeaderWritten = true
	w.Done = true
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *GBonResponseWriter) Write(content []byte) (int, error) {
	w.Done = true
	return w.ResponseWriter.Write(content)
}
