package main

import (
	"net/http"
	"notchman8600/authentication-provider/infra"
)

func main() {
	router := infra.NewRouter()
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", infra.IndexHandler)
	http.ListenAndServe(":8080", router)
}
