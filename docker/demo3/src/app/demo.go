package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}
	clientip := getCurrentIP(r)
	fmt.Println(r.RemoteAddr)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientip)
	fmt.Fprintln(w, "hello world")
}
func healthz(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "working") }

func handlerOne(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		writer.Header().Set(k, v[0])
	}
	fmt.Println("the request body is", request.Body)
	//set response header
	writer.Header().Set("content-Type", "application/json")
	writer.WriteHeader(200)
	fmt.Fprintln(writer, "this is hello response body")
}
func getCurrentIP(r *http.Request) string {
	// 这里也可以通过X-Forwarded-For请求头的第一个值作为用户的ip
	//但是要注意的是这两个请求头代表的ip都有可能是伪造的
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
func main() {
	mux := http.NewServeMux()
	// 06. debug
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/you", handlerOne)
	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	}
}
