package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type class struct {
	Index string `json:"index"`
}

func getVal(id string, jsonBlob []byte) (*class, error) {
	var classes []class
	json.Unmarshal(jsonBlob, &classes)
	if classes == nil {
		return nil, errors.New("no json data")
	}

	for _, val := range classes {
		if val.Index == id {
			return &val, nil
		}
	}

	return nil, errors.New("id not found")
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func classHandler(w http.ResponseWriter, req *http.Request) {
	// read file
	bytes, err := os.ReadFile(fmt.Sprintf("data/%s/5e-SRD-Classes.json", req.PathValue("year")))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	class, err := getVal(req.PathValue("id"), bytes)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	classBytes, err := json.Marshal(class)
	if err != nil {
		w.Write([]byte("error marshalling class"))
		return
	}

	w.Write(classBytes)
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

	mux.HandleFunc("/{year}/classes/{id}", classHandler)

	http.ListenAndServe(":8080", logger(mux))
}
