package infra

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"notchman8600/authentication-provider/interfaces/controller"
	"os"
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
var Templates = make(map[string]*template.Template)

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

func authHandler(controller *controller.AuthController) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := controller.Auth(r)
		if err != nil {
			//TODO エラーレスポンスの作成
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error! \r\n")
			return
		}
		// CookieにセッションIDをセット
		cookie := &http.Cookie{
			Name:  "session",
			Value: session.Id,
		}
		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "create session! \r\n")

	})
}

func authCheckHandler(controller *controller.AuthController) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.AuthCheck(w, r)
	})
}

func tokenHandler(controller *controller.AuthController) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.Token(w, r)
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
	url := os.Getenv("SQL_URL")
	sqlHandler := NewDB("mysql", url)
	authController := controller.NewAuthController(sqlHandler)

	router.Methods(http.MethodGet).Handler("/", indexHandler())
	router.Methods(http.MethodGet).Handler("/sample", sampleHandler())
	router.Methods(http.MethodPost).Handler("/auth", authCheckHandler(authController))
	router.Methods(http.MethodGet).Handler("/auth", authHandler(authController))
	router.Methods(http.MethodGet).Handler("/token", tokenHandler(authController))
	return router
}
