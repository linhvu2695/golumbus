package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	logger *log.Logger
}

func NewGoodbye(log *log.Logger) *Goodbye {
	return &Goodbye{log}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte ("Adios!!"))
}