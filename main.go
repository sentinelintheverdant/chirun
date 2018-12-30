package main

import (
	. "fmt"

	"net/http"

	"github.com/go-chi/chi"
)

const (
	port = "80"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Println("Hit the endpoint")
		w.Write([]byte("welcome"))
	})

	Println("Server Listening:")
	Printf("    ==> localhost:%v\n", port)
	http.ListenAndServe(":" + port, r)
}