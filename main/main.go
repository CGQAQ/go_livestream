package main

import (
	"io/ioutil"
	"net/http"
	livego "github.com/gwuhaolin/livego"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: MyHandler{},
	}
	_ = server.ListenAndServe()

}

type MyHandler struct {
}

func (handler MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/index.html" {
		res, _ := ioutil.ReadFile("frontend/index.html")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		_, _ = w.Write(res)
	}
}