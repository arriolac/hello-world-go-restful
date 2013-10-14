package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"io"
	"net/http"
)

// This example shows the minimal code needed to get a restful.WebService working.
//
// GET http://localhost:8080/hello

func main() {
	fmt.Println("Listening...")
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	http.ListenAndServe(":8080", nil)
}

func hello(req *restful.Request, resp *restful.Response) {
	fmt.Println("Received request: ", req)
	io.WriteString(resp, "world")
}
