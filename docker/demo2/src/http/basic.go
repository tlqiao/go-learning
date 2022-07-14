package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HandlerOne)
	http.ListenAndServe(":9099", nil)
}

func HandlerOne(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		writer.Header().Set(k, v[0])
	}
	fmt.Println("the request body is", request.Body)
	//set response header
	writer.Header().Set("content-Type", "application/json")
	writer.WriteHeader(200)
	fmt.Fprintln(writer, "this is hello response body")
}
