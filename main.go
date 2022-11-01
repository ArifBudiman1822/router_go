package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	server := http.Server{
		Handler: r,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
