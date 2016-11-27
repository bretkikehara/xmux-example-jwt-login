package main

import (
	"log"
	"net/http"

	"github.com/rs/xhandler"
	"golang.org/x/net/context"
)

type ExampleLogger struct {
	next xhandler.HandlerC
}

func (h ExampleLogger) ServeHTTPC(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	log.Printf("URI: %s", r.RequestURI)
	log.Printf("\tAccept: %s", r.Header.Get("Accept"))
	log.Printf("\tContent-Type: %s", r.Header.Get("Content-Type"))
	h.next.ServeHTTPC(ctx, w, r)
}
