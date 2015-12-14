package controller
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"models"
	"encoding/json"
	"fmt"
)

type UserController struct {}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	u := models.User{
		Name: "Sunheng",
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Fprintf(w, "404 Not Found%s", id)
	}
	setReturnHeaders(w)
	fmt.Fprintf(w, "%s", uj)
}

func setReturnHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}