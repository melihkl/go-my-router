package router

import (
	"net/http"
	"strings"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, map[string]string)

type Router struct {
	routes map[string]map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (r *Router) GET(path string, handler HandlerFunc) {
	r.addRoute(http.MethodGet, path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.addRoute(http.MethodPost, path, handler)
}

func (r *Router) addRoute(method, path string, handler HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	methodRoutes := r.routes[req.Method]
	if methodRoutes == nil {
		http.NotFound(w, req)
		return
	}

	for routePath, handler := range methodRoutes {
		params, ok := matchPath(routePath, req.URL.Path)
		if ok {
			handler(w, req, params)
			return
		}
	}

	http.NotFound(w, req)
}

func matchPath(routePath, requestPath string) (map[string]string, bool) {
	routeParts := strings.Split(routePath, "/")
	requestParts := strings.Split(requestPath, "/")

	if len(routeParts) != len(requestParts) {
		return nil, false
	}

	params := make(map[string]string)

	for i := range routeParts {
		if strings.HasPrefix(routeParts[i], ":") {
			paramName := routeParts[i][1:]
			params[paramName] = requestParts[i]
		} else if routeParts[i] != requestParts[i] {
			return nil, false
		}
	}
	return params, true
}
