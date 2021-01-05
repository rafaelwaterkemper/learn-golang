package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

// func main() {
// 	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	})

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello from base path"))
// 	})

// 	log.Fatal("Serving", http.ListenAndServe(":8080", nil))
// }

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        &GlobalHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

type GlobalHandler struct {
}

func (*GlobalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q from httpServer configured", html.EscapeString(r.URL.Path))
}
