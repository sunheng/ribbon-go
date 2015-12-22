package controller
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"models"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	var oid bson.ObjectId
	if bson.IsObjectIdHex(id) {
		oid = bson.ObjectIdHex(id)
	} else {
		setNotFoundHeaders(w)
		return
	}

	var user models.User
	if err := uc.session.DB("ribbon").C("users").FindId(oid).One(&user); err != nil {
		setNotFoundHeaders(w)
		return
	}

	writeResponse(w, 200, user)
}

func (uc *UserController) AddUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	user.ID = bson.NewObjectId()
	if err := uc.session.DB("ribbon").C("users").Insert(user); err != nil {
		setNotFoundHeaders(w)
		return
	}

	writeResponse(w, 201, user)
}

func writeResponse(w http.ResponseWriter, code int, output interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if jsonOutput, err := json.Marshal(output); err != nil {
		setNotFoundHeaders(w)
	} else {
		fmt.Fprintf(w, "%s", jsonOutput)
	}
}

func setNotFoundHeaders(w http.ResponseWriter) {
	w.WriteHeader(404)
}