package main

import (
	"net/http"
	"os"
	"fmt"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, r.Header.Get("X-Forwarded-For"))
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}

