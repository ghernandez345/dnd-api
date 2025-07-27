package main

import "fmt"
import "net/http"

func testHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hellow World!")
}

func main () {
	mux := http.NewServeMux()

	mux.HandleFunc("/", testHandler)

	http.ListenAndServe(":8080", mux)

}
