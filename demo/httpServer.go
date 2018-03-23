package main

import (
	"fmt"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	println(r.Method)
	if r.Method != "GET" {
		fmt.Fprintf(w, "GET method not support\n")
	}

	fmt.Fprintf(w, "Hello there!\n")
}

func main() {
	http.HandleFunc("/hello", myHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8081", nil))
}
