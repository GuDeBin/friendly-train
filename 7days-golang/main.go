package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>Hello World golang</h1>")
	})

	// 打印请求链接
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.Path)
	})

	http.ListenAndServe(":3000", nil)
}
