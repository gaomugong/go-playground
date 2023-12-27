package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	restful "github.com/emicklei/go-restful/v3"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResource struct {
	users map[string]User
}

func (us *UserResource) createUser(request *restful.Request, response *restful.Response) {
	u := User{ID: fmt.Sprintf("%d", time.Now().Unix())}
	err := request.ReadEntity(&u)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	us.users[u.ID] = u
	response.WriteHeaderAndEntity(http.StatusCreated, u)
}

func (us *UserResource) findAllUser(request *restful.Request, response *restful.Response) {
	users := []User{}
	for _, u := range us.users {
		users = append(users, u)
	}
	response.WriteEntity(users)
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}

func main() {
	// us := new(UserResource)
	// us.users = make(map[string]User)
	us := UserResource{users: map[string]User{
		"zhangsan": User{"zhangsan", "张三", 10},
		"lisi":     User{"lisi", "李四", 11},
	}}

	container := restful.NewContainer()

	//  create webservice
	ws := new(restful.WebService)
	ws.Path("/api").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/hello").To(hello))
	ws.Route(ws.POST("/users").To(us.createUser))
	ws.Route(ws.GET("/users").To(us.findAllUser))
	// restful.Add(ws)
	container.Add(ws)
	log.Println("start server :8080")
	if err := http.ListenAndServe(":8080", container); err != nil {
		log.Fatalf("start server error: %s", err)
	}
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
