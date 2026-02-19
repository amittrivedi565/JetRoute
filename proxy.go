package main

import (
	"fmt"
	"net/http"
)

func proxy() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Request Received]", r.URL.Path)
		fmt.Fprintln(w, "JetRoute [Request Received]")
	})

	fmt.Println("JetRoute [Server Started] on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
