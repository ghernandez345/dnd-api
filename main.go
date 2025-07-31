package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func classHandler(w http.ResponseWriter, req *http.Request) {
	// read file
	bytes, err := os.ReadFile(fmt.Sprintf("data/%s/5e-SRD-%s", req.PathValue("year"), req.PathValue("id")))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(bytes)
	// get class id from req path and get obj val by this key

	// create struct and add it to w writer as json
}

func logger(next http.Handler) http.Handler {
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

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", testHandler)
	mux.HandleFunc("/test2", test2Handler)
	mux.HandleFunc("/{year}/classes/{id}", classHandler)

	http.ListenAndServe(":8080", logger(mux))
}
