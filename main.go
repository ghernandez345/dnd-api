package main

import (
	"net/http"
	"log/slog"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		slog.Info("request", "method", req.Method, "url", req.URL)
		next.ServeHTTP(w, req)
	})
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World!"))
}

func test2Handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello two!"))
}


func main () {
	mux := http.NewServeMux()

	mux.HandleFunc("/", testHandler)
	mux.HandleFunc("/test2", test2Handler)

	http.ListenAndServe(":8080", loggingMiddleware(mux))
}
