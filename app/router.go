package app

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Router represents interface that each router should implement
type Router interface {
	http.Handler

	Get(pattern string, handlerFn http.HandlerFunc)
	Post(pattern string, handlerFn http.HandlerFunc)
	Put(pattern string, handlerFn http.HandlerFunc)
	Patch(pattern string, handlerFn http.HandlerFunc)
	Delete(pattern string, handlerFn http.HandlerFunc)
	Options(pattern string, handlerFn http.HandlerFunc)
	Use(middleware func(http.Handler) http.Handler)
	With(middleware func(http.Handler) http.Handler) Router
	Route(pattern string, fn func(r Router)) Router
	Mount(pattern string, h http.Handler)
}

// Controller represents interface that each controller should implement
type Controller interface {
	Register(r Router)
}

// ChiRouter is a wrapper arround chi.Router
type ChiRouter struct {
	chiRouter chi.Router
}

// NewChiRouter creates new router
func NewChiRouter() Router {
	return &ChiRouter{chi.NewRouter()}
}

func (router *ChiRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.chiRouter.ServeHTTP(w, r)
}

func (router *ChiRouter) Get(pattern string, handlerFn http.HandlerFunc) {
	router.chiRouter.Get(pattern, handlerFn)
}

func (router *ChiRouter) Post(pattern string, handlerFn http.HandlerFunc) {
	router.chiRouter.Post(pattern, handlerFn)
}

func (router *ChiRouter) Put(pattern string, handlerFn http.HandlerFunc) {
	router.chiRouter.Put(pattern, handlerFn)
}

func (router *ChiRouter) Delete(pattern string, handlerFn http.HandlerFunc) {
	router.chiRouter.Delete(pattern, handlerFn)
}

func (router *ChiRouter) Patch(pattern string, handlerFn http.HandlerFunc) {
	router.chiRouter.Patch(pattern, handlerFn)
}

func (router *ChiRouter) Options(pattern string, handlerFn http.HandlerFunc) {
	router.chiRouter.Options(pattern, handlerFn)
}

func (router *ChiRouter) Use(middleware func(http.Handler) http.Handler) {
	router.chiRouter.Use(middleware)
}

func (router *ChiRouter) With(middleware func(http.Handler) http.Handler) Router {
	subChiRouter := router.chiRouter.With(middleware)
	return &ChiRouter{subChiRouter}
}

func (router *ChiRouter) Route(pattern string, fn func(r Router)) Router {
	subRouter := NewChiRouter()
	if fn != nil {
		fn(subRouter)
	}
	router.Mount(pattern, subRouter)
	return subRouter
}

func (router *ChiRouter) Mount(pattern string, h http.Handler) {
	router.chiRouter.Mount(pattern, h)
}

// URLParam returns the url parameter as a string
// from a http.Request object.
func URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

// URLParamInt returns the url parameter as an integer
// from a http.Request object.
func URLParamInt(r *http.Request, key string) (int, error) {
	return strconv.Atoi(URLParam(r, key))
}
