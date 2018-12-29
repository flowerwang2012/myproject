package main

import "net/http"

func main() {
	//http.Handle("/hello", )
	http.HandleFunc("/world", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("world"))
	})
	http.ListenAndServe(":8080", nil)
}