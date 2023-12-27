package main

import (
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

func main() {
	container := restful.NewContainer()

	//  create webservice
	ws := new(restful.WebService)
	ws.Path("/api").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/hello").To(hello))

	// restful.Add(ws)
	container.Add(ws)
	log.Println("start server :8080")
	if err := http.ListenAndServe(":8080", container); err != nil {
		log.Fatalf("start server error: %s", err)
	}
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
