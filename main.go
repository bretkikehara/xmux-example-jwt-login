package main

import (
	"fmt"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
	"log"
	"net/http"
	"time"
)

func main() {
	c := xhandler.Chain{}

	// Add close notifier handler so context is cancelled when the client closes
	// the connection
	c.UseC(xhandler.CloseHandler)

	// Add timeout handler
	c.UseC(xhandler.TimeoutHandler(2 * time.Second))

	c.UseC(func(next xhandler.HandlerC) xhandler.HandlerC {
		return ExampleLogger{next: next}
	})

	mux := xmux.New()

	mux.Handle("GET", "/login.html", http.FileServer(http.Dir("./static")))

	mux.POST("/auth", xhandler.HandlerFuncC(AuthenticateHandler))
	// mux.GET("/token", xhandler.HandlerFuncC(token))

	port := getPort()

	log.Printf("Running server on localhost:%d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), c.Handler(mux)); err != nil {
		log.Fatal(err)
	}
}
