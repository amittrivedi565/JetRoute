package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// -> http://jetroute/order-service/users/123
// -> http://localhost:service-port/users/123
func router(w http.ResponseWriter, r *http.Request) {

	service, endpoint := extractService(r.URL.Path)

	cfg, ok := Config[service]
	if !ok {
		http.Error(w, "service not found", http.StatusNotFound)
		return
	}
	target, _ := url.Parse(fmt.Sprintf("http://%s:%d", cfg.Host, cfg.Port))

	r.URL.Path = endpoint
	r.Host = target.Host

	httputil.NewSingleHostReverseProxy(target).ServeHTTP(w, r)
}

func extractService(path string) (string, string) {

	path = strings.Trim(path, "/")
	parts := strings.SplitN(path, "/", 2)

	service := parts[0]

	endpoint := "/"
	if len(parts) > 1 {
		endpoint += parts[1]
	}

	return service, endpoint
}
