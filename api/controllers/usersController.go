package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	//fmt.Println(user)

	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	res := user.Create()
	fmt.Println(res)
	u.Respond(w, res)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
	}

	res := models.Login(user.Name, user.Password)

	u.Respond(w, res)
}

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	// user := r.Context().Value("user").(uint)
	params := mux.Vars(r)
	id := params["id"]
	data := models.GetUser(id)

	res := u.Message(true, "success")

	res["data"] = data

	u.Respond(w, res)
}

var GetUsers = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetUsers()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}

var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	user := &models.User{}

	err2 := json.NewDecoder(r.Body).Decode(user)
	if err2 != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	res := user.UpdateUser(id)
	res["data"] = id
	u.Respond(w, res)
}

var UpdateUserPassword = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	user := &models.User{}

	err2 := json.NewDecoder(r.Body).Decode(user)
	if err2 != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	res := models.UpdateUserPassword(id, user.Password)
	res["data"] = id
	u.Respond(w, res)
}

var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res := models.DeleteUser(id)
	res["data"] = id
	u.Respond(w, res)
}
