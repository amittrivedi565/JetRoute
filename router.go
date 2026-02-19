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

		private := false

		if strings.HasSuffix(route.Path, "/*") {
			prefix := strings.TrimSuffix(route.Path, "/*")
			private = strings.HasPrefix(endpoint, prefix)
		} else {
			private = endpoint == route.Path
		}

		if private {
			if !authorize(r, cfg) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			break
		}
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

func authorize(r *http.Request, cfg *ServiceConfig) bool {

	authURL := fmt.Sprintf(
		"http://%s:%d%s",
		cfg.Auth.Host,
		cfg.Auth.Port,
		cfg.Auth.Path,
	)

	req, err := http.NewRequest("GET", authURL, nil)
	if err != nil {
		return false
	}

	req.Header = r.Header.Clone()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
