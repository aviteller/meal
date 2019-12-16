package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	u "../utils"
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
