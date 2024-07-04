package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexhander)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe("9999", nil)
}

func indexhander(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, req.URL)
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, req.URL)
}
