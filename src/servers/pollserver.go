package main
import (
	"fmt"
	"time"
	"models"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

func main() {
	reminderColl := mongoSession().DB("ribbon").C("reminders")
//	{dueTime: {$lte: ISODate("2015-12-21T12:00:00.000Z")}}
 	var reminders []*models.Reminder
	for {
		q := bson.M{
			"dueTime": bson.M{
				"$lte": time.Now(),
			},
		}
		reminderColl.Find(q).All(&reminders)
		fmt.Println(len(reminders))
		time.Sleep(time.Duration(15) * time.Second)
		fmt.Print(".")
	}
}

func mongoSession() *mgo.Session {
	// Connect to our mongo

	s, err := mgo.Dial("mongodb://localhost:8080")
	if err != nil {
		panic(err)
	}

	return s
}