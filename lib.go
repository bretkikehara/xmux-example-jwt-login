package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func ParseRequestBodyAsJson(r *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(obj)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

func getPort() int {
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		return port
	}
	return 8080
}
