package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nbari/violetear"
	"github.com/swisscom/cf-sample-app-go-dep/app/hello"
)

var version string

func main() {
	version = os.Getenv("VERSION")

	router := violetear.New()
	router.HandleFunc("/health", health)
	router.HandleFunc("/", hello.Handler)

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   7 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(srv.ListenAndServe())
}
