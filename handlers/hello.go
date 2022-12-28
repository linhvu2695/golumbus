package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Hello World")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	log.Printf("Data: %s\n", data)

	fmt.Fprintf(w, "Server response: Hello %s\n", data)
}