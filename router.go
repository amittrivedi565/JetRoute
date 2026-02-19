package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func router(w http.ResponseWriter, r *http.Request) {
	service, endpoint := extractService(r.URL.Path)

	cfg, ok := Config[service]
	if !ok {
		http.Error(w, "service not found", http.StatusNotFound)
		return
	}
	for _, route := range cfg.PrivateRoutes {
		if strings.HasSuffix(route.Path, "/*") {
			prefix := strings.TrimPrefix(route.Path, "/*")
			if strings.HasSuffix(endpoint, prefix) {
				// call the auth service
			}
		} else {
			target, _ := url.Parse(fmt.Sprintf("http://%s:%d", cfg.Host, cfg.Port))

			r.URL.Path = endpoint
			r.Host = target.Host

			httputil.NewSingleHostReverseProxy(target).ServeHTTP(w, r)
		}
	}
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
