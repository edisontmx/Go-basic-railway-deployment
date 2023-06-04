package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world: %s\n", r.URL.Path)
	})
	puerto := os.Getenv("PORT")
	if puerto == "" {
		puerto = "8080"
	}
	http.ListenAndServe(":"+puerto, nil)

}
