package main

import "fmt"
import "net/http"

func main () {
	testHandler := func (w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hellow World!")
	}

	http.HandleFunc("/", testHandler)
	http.ListenAndServe(":8080", nil);
}
