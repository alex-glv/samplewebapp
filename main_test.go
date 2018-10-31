package main

import (
	"net/http"
	"net/url"
	"testing"
)

type DummyWriter struct {
	Output     []byte
	StatusCode int
}

func (dw *DummyWriter) Header() http.Header {
	return http.Header{}
}

func (dw *DummyWriter) Write(input []byte) (int, error) {
	dw.Output = input
	return 0, nil
}

func (dw *DummyWriter) WriteHeader(statusCode int) {
	dw.StatusCode = statusCode
}

func GetRequest(path string) *http.Request {
	return &http.Request{
		URL: &url.URL{
			Path: path,
		},
	}
}

func TestResponse(t *testing.T) {
	server := &SampleHandler{}
	w := &DummyWriter{}
	r := GetRequest("/test")
	server.ServeHTTP(w, r)
	expected := "Hi there, got 1 requests, I love test!"
	if string(w.Output) != expected {
		t.Error("Expected ", expected, " got ", string(w.Output))
	}
}
