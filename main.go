package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	requestCounter int
)

type SampleHandler struct{}

func (h *SampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestCounter++
	logrus.Infoln("Request num:", requestCounter)
	fmt.Fprintf(w, "Hi there, got %d requests, I love %s!", requestCounter, r.URL.Path[1:])
}

func main() {
	requestCounter = 0
	server := &http.Server{
		Handler: &SampleHandler{},
		Addr:    ":8080",
	}
	logrus.Fatal(server.ListenAndServe())
}
