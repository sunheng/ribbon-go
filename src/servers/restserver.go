package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"controller"
	"gopkg.in/mgo.v2"
)

func main() {
	router := httprouter.New()
	userController := controller.NewUserController(getSession())
	reminderController := controller.NewReminderController(getSession())

	router.GET("/user/:id", userController.GetUser)
	router.POST("/user", userController.AddUser)

	router.POST("/reminder", reminderController.AddReminder)

	http.ListenAndServe("localhost:3000", router)
}

func getSession() *mgo.Session {
	// Connect to our mongo

	s, err := mgo.Dial("mongodb://localhost:8080")
	if err != nil {
		panic(err)
	}

	return s
}