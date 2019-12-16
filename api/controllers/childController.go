package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	u "../utils"
)

var CreateChild = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	child := &models.Child{}

	err := json.NewDecoder(r.Body).Decode(child)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	child.UserID = int(user)
	res := child.Create()
	u.Respond(w, res)
}

var GetChildren = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)

	data := models.GetChildrenByUser(int(user))
	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)
}
