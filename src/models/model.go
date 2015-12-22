package models
import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	User struct {
		ID 		bson.ObjectId	`json:"id" bson:"_id"`
		Name 	string			`json:"name" bson:"name"`
	}

	Reminder struct {
		ID		bson.ObjectId	`json:"id" bson:"_id"`
		Title	string			`json:"title" bson:"title"`
		UserID 	bson.ObjectId	`json:"userId" bson:"userId"`
		DueTime	time.Time		`json:"dueTime" bson:"dueTime"`
	}
)