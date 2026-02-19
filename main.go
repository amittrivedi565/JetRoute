package main

import "log"

func main() {

	err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	proxy()
}
