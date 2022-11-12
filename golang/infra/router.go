package infra

import (
	"errors"
	"fmt"
	"net/http"
)

type Router struct {
	tree                  *tree
	NotFoundHandler       http.Handler
	MethodNotAllowHandler http.Handler
}

type route struct {
	methods []string
	path    string
	handler http.Handler
}

var (
	tmpRoute            = &route{}
	ErrNotFound         = errors.New("no matching route was found")
	ErrMethodNotAllowed = errors.New("method is not allowed")
)

func (r *Router) Handle() {
	r.tree.Insert(tmpRoute.methods, tmpRoute.path, tmpRoute.handler)
	tmpRoute = &route{}
}

func (r *Router) Methods(methods ...string) *Router {
	tmpRoute.methods = append(tmpRoute.methods, methods...)
	return r
}

func (r *Router) Handler(path string, handler http.Handler) {
	tmpRoute.handler = handler
	tmpRoute.path = path
	r.Handle()
}

func handleErr(err error) int {
	var status int
	switch err {
	case ErrMethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case ErrNotFound:
		status = http.StatusNotFound
	}
	return status
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	//該当のハンドラーを見つける
	result, err := r.tree.Search(method, path)
	if err != nil {
		status := handleErr(err)
		w.WriteHeader(status)
		return
	}
	h := result.actions.handler
	h.ServeHTTP(w, req)
}

func indexHandler() http.Handler {
	// switching http method
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /\r\n")
	})
}

func authHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func sampleHandler() http.Handler {
	// switching http method
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! \r\n")
	})
}

func NewRouter() *Router {
	router := &Router{
		tree: NewTree(),
	}

	router.Methods(http.MethodGet).Handler("/", indexHandler())
	router.Methods(http.MethodGet).Handler("/sample", sampleHandler())
	return router
}
