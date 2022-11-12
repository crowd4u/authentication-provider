package main

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"notchman8600/authentication-provider/infra"
)

func base64URLEncode(verifier string) string {
	hash := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", infra.IndexHandler)
	http.ListenAndServe(":8080", mux)
}
