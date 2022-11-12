package infra

import (
	"errors"
	"net/http"
)

type Router struct {
	tree                  *tree
	NotFoundHandler       http.Handler
	MethodNotAllowHandler http.Handler
}

type route struct {
	method  []string
	path    string
	handler http.Handler
}

var (
	tmpRoute            = &route{}
	ErrNotFound         = errors.New("no matching route was found")
	ErrMethodNotAllowed = errors.New("Method is not allowed")
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// switching http method
	switch r.Method {
	case http.MethodGet:
		// TODO
	case http.MethodPost:
		// TODO

	}
}
