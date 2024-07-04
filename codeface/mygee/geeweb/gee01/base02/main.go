package main

import (
	"fmt"
	"net/http"
)

type Engine struct {
}

func (e Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		_, err := fmt.Fprintf(writer, "request.URL.Path=%q/n", request.URL.Path)
		if err != nil {
			return
		}
	case "/hello":
		for k, v := range request.Header {
			_, err := fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
			if err != nil {
				return
			}
		}

	default:
		_, err := fmt.Fprintf(writer, "NOT FIND")
		if err != nil {
			return
		}

	}
}

func main() {
	engine := new(Engine)
	err := http.ListenAndServe(":9999", engine)
	if err != nil {
		return
	}
}
