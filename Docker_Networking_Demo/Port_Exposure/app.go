package main

import (
	"fmt"
	"net/http"
)

func main() {
	finish := make(chan bool)

	server1 := http.NewServeMux()
	server1.HandleFunc("/", HandlerV1)

	server2 := http.NewServeMux()
	server2.HandleFunc("/v2", HandlerV2)
	server2.HandleFunc("/v1", HandlerV1)

	go func() {
		fmt.Println("Starting server on port 8001...")
		err := http.ListenAndServe(":8001", server1)
		if err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		fmt.Println("Starting server on port 8002...")
		err := http.ListenAndServe(":8002", server2)
		if err != nil {
			fmt.Println(err)
		}
	}()

	<-finish
}

func HandlerV1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func HandlerV2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greetings!")
}
