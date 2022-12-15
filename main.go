package main

import (
	"fmt"
	"io"
	"log"
	h "net/http"
)

func main() {

	f1 := func(w h.ResponseWriter, _ *h.Request) {
		io.WriteString(w, "birinci func \n ")
	}

	f2 := func(w h.ResponseWriter, _ *h.Request) {
		io.WriteString(w, "ikinci func \n ")
	}

	h.HandleFunc("/", f1)
	h.HandleFunc("/f2", f2)
	h.HandleFunc("/f3", f3)

	log.Fatal(h.ListenAndServe(":8081", nil))
}

func f3(w h.ResponseWriter, r *h.Request) {
	// io.WriteString(w, "3 func")
	fmt.Printf("r.Method: %v\n", r.Method)
	io.WriteString(w, r.RequestURI)
}
