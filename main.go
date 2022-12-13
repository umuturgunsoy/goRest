package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	f1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "birinci func \n ")
	}

	f2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "ikinci func \n ")
	}

	http.HandleFunc("/", f1)
	http.HandleFunc("/f2", f2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
