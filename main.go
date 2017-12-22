package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController struct {
	Controller
	registry *registry
}

func main() {
	registry := setup()

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", registry.router)
}

func configureRoutes(registry *registry)  {
	c := &UserController{registry: registry}

	registry.router.GET("/api/user", c.Perform(c.getList))
	registry.router.POST("/api/user", c.Perform(c.AddUser))

}

func (c *Controller) Index(rw http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	return nil
}
