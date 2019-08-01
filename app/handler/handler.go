package handler

import (
	"fmt"
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type Router struct {
	routes []*route
}

func (t *Router) handler(pattern *regexp.Regexp, handler http.Handler) {
	t.routes = append(t.routes, &route{pattern, handler})
}

func (t *Router) handleFunc(pattern *regexp.Regexp, handlerFunc func(http.ResponseWriter, *http.Request)) {
	t.handler(pattern, http.HandlerFunc(handlerFunc))
}

func (t *Router) HandleFunc(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) error {
	if expr, err := regexp.Compile(pattern); err == nil{
		t.handleFunc(expr, handlerFunc)
	}else{
		return err
	}
	return nil
}

func (t Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range t.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}

	http.NotFound(w, r)
}

func GetAllCoupons(writer http.ResponseWriter, requestPtr *http.Request) {
	fmt.Fprintln(writer, "hello")
}
