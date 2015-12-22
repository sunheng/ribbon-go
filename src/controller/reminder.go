package controller
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"models"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ReminderController struct {
	session *mgo.Session
}

func NewReminderController(s *mgo.Session) *ReminderController {
	return &ReminderController{s}
}

func (rc *ReminderController) AddReminder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var reminder models.Reminder

	json.NewDecoder(r.Body).Decode(&reminder)
	reminder.ID = bson.NewObjectId()
	if err := rc.session.DB("ribbon").C("reminders").Insert(reminder); err != nil {
		setNotFoundHeaders(w)
		return
	}

	writeResponse(w, 201, reminder)
}