package main

import (
	"fmt"
	"net/http"
)

const (
	TOKEN = "123456" // this is the password on my luggage
)

func main() {
	http.HandleFunc("/", handler)

	//listen on port 9000
	port := ":9000"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic("Error starting the http server")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Token: "+TOKEN)
	if err != nil {
		panic("Error writing to the response writer")
	}
}
