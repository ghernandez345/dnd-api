package main

import (
	"fmt"
	"net/http"
	"log/slog"
)

func testHandler(w http.ResponseWriter, req *http.Request) {
	slog.Info("request", "method", req.Method, "url", req.URL)
	fmt.Fprint(w, "Hellow World!")
}

func main () {
	mux := http.NewServeMux()

	mux.HandleFunc("/", testHandler)

	http.ListenAndServe(":8080", mux)
}
