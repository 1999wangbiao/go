package main

import (
	"bytes"
	"fmt"

	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var resp *http.Response
	header := make(map[string]string)
	header["Accept"] = "*/*"
	url := "http://zwgk.f3322.net:380/api/v1/pYp7CLul94FyTJOBqoDq/rpc"
	msg := `{"method": "getInfo", "params":{}}`
	fmt.Println(msg)
	fmt.Println(header)

	resp, err := HttpRequest("Post", url, []byte(msg), header)
	if err != nil {
		fmt.Println("aaa")
	}

	fmt.Println("aaa")
	a, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("resq.Body:", string(a))
}

func HttpRequest(method string, url string, msg []byte, header map[string]string) (resp *http.Response, err error) {
	fmt.Println("lib/util/http.go/HttpRequest")
	var req *http.Request
	fmt.Println("url", url)

	req, err = http.NewRequest(strings.ToUpper(method), url, bytes.NewReader(msg))
	if err != nil {
		fmt.Println("sfs")
		return
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	client := http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("sss")
		return
	}

	return
}
