package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", health)

	log.Printf("service is listening at port %s ", "3000")
	http.ListenAndServe(":3000", r)

}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("service is running"))
}
