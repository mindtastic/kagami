package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":7666", "address to listen on")

func main() {
	flag.Parse()

	log.Printf("listening on address %q", *addr)
	log.Fatal(http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("new request from %q", r.RemoteAddr)
		for k, v := range r.Header {
			w.Write([]byte(fmt.Sprintf("%s: %s\n", k, v)))
		}
	})))
}
