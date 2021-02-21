package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/zozoee27/cookbook/backend/entity"
	"github.com/zozoee27/cookbook/backend/service"
	"github.com/zozoee27/cookbook/backend/util"
)

func MakeUserHandlers(r *mux.Router, s *service.User) {
	r.Handle("/account/register", handlers.ContentTypeHandler(
		handlers.LoggingHandler(os.Stdout,
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				registerAccount(w, r, s)
			})), "application/json")).Methods("POST")
}

func registerAccount(w http.ResponseWriter, r *http.Request, s *service.User) {
	log.Print("Recieved /account/register command")

	var user entity.User

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)

	if err != nil {
		log.Print("AccountManager::RegisterAccount: Cannot unmarshal json")
		return
	}

	err = s.CreateUser(&user)
	if err != nil {
		log.Print("Could not add user to collection: ", err)
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondWithCode(w, http.StatusCreated)
}
