package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type LoginJSON struct {
	Username, Passwd string
}

func Authenticate(l LoginJSON) bool {
	log.Printf("%s:%s", l.Username, l.Passwd)
	return l.Username == "test" && l.Passwd == "test"
}

func AuthenticateHandler(_ context.Context, w http.ResponseWriter, r *http.Request) {
	var login LoginJSON
	if err := ParseRequestBodyAsJson(r, &login); err == nil && Authenticate(login) {
		fmt.Fprint(w, "Authorized\n")
	} else {
		fmt.Fprint(w, "Not authorized\n")
	}
}
