package router

import (
	"net/http"
	"strings"
)

// HandlerFunc: isteğe göre parametreleri de alan handler tipi
type HandlerFunc func(http.ResponseWriter, *http.Request, map[string]string)

// Router yapısı: HTTP methoduna göre handler'ları tutar
type Router struct {
	routes map[string]map[string]HandlerFunc
}

// Yeni bir Router oluşturur
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

// GET için route ekler
func (r *Router) GET(path string, handler HandlerFunc) {
	r.addRoute(http.MethodGet, path, handler)
}

// POST için route ekler
func (r *Router) POST(path string, handler HandlerFunc) {
	r.addRoute(http.MethodPost, path, handler)
}

// ✅ PUT için route ekler
func (r *Router) PUT(path string, handler HandlerFunc) {
	r.addRoute(http.MethodPut, path, handler)
}

// ✅ DELETE için route ekler
func (r *Router) DELETE(path string, handler HandlerFunc) {
	r.addRoute(http.MethodDelete, path, handler)
}

// Ortak route ekleme fonksiyonu
func (r *Router) addRoute(method, path string, handler HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][path] = handler
}

// HTTP isteklerini karşılar
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	methodRoutes := r.routes[req.Method]
	if methodRoutes == nil {
		http.NotFound(w, req)
		return
	}

	for routePath, handler := range methodRoutes {
		if params, ok := matchPath(routePath, req.URL.Path); ok {
			handler(w, req, params)
			return
		}
	}

	http.NotFound(w, req)
}

// /users/:id gibi yolları /users/123 ile eşleştiren fonksiyon
func matchPath(pattern, actual string) (map[string]string, bool) {
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	actualParts := strings.Split(strings.Trim(actual, "/"), "/")

	if len(patternParts) != len(actualParts) {
		return nil, false
	}

	params := make(map[string]string)
	for i := 0; i < len(patternParts); i++ {
		if strings.HasPrefix(patternParts[i], ":") {
			paramName := patternParts[i][1:]
			params[paramName] = actualParts[i]
		} else if patternParts[i] != actualParts[i] {
			return nil, false
		}
	}

	return params, true
}
