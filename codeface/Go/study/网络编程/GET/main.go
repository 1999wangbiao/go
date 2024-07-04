package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	resp,err:=http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Println("get failed ,err:\n",err)
		return
	}
	defer resp.Body.Close()
	body, err:= ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read form resp.Body failed,err:\n",err)
		return
	}
	fmt.Println(string(body))
}
