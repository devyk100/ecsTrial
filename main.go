package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Post("/", func(res http.ResponseWriter, req *http.Request) {
		_, err := res.Write([]byte("Hello World"))
		if err != nil {
			log.Println("Error writing response:", err)
			return
		}
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
	fmt.Println("Hello World")
}
