package main

import (
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Path    string `json:"path"`
}

func proxy() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("[Request Received] %s\n", r.URL.Path)
		router(w, r)
	})

	fmt.Println("[Server Started Listening on port 8080]")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
