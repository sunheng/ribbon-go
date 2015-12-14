package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"controller"
)

func main() {
	router := httprouter.New()
	userController := controller.NewUserController()
	router.GET("/user/:id", userController.GetUser)

	http.ListenAndServe("localhost:3000", router)
}