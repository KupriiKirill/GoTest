package handler

import (
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
	method  string
}

const (
	// GET is a const for httpmethod
	GET     = "GET"
	// PUT is a const for httpmethod
	PUT     = "PUT"
	// DELETE is a const for httpmethod
	DELETE  = "DELETE"
	// HEAD is a const for httpmethod
	HEAD 	= "HEAD"
)

//Router is a basic request handler with only HandleFunc function
type Router struct {
	routes []*route
}

func (t *Router) handler(pattern *regexp.Regexp, handler http.Handler, method string) {
	t.routes = append(t.routes, &route{pattern, handler, method})
}

func (t *Router) handleFunc(pattern *regexp.Regexp, handlerFunc func(http.ResponseWriter, *http.Request), method string) {
	t.handler(pattern, http.HandlerFunc(handlerFunc), method)
}

//HandleFunc allows to set a specific route handler
func (t *Router) HandleFunc(pattern string, handlerFunc func(http.ResponseWriter, *http.Request), method string) error {
	if expr, err := regexp.Compile(pattern); err == nil {
		t.handleFunc(expr, handlerFunc, method)
	} else {
		return err
	}
	return nil
}

func (t Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handled := false
	found := false

	for _, route := range t.routes {
		if route.pattern.MatchString(r.URL.Path) {
			found = true
			if r.Method == route.method {
				handled = true
				route.handler.ServeHTTP(w, r)
				return
			}
		}
	}

	if !found {
		http.NotFound(w, r)
		return
	}

	if !handled {
		http.Error(w, "Method is not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
