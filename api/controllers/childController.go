package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
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

var GetChildrenByUser = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)

	data := models.GetChildrenByUser(int(user))
	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)
}

var GetChildren = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetChildren()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}

var UpdateChild = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	note := &models.Child{}

	err2 := json.NewDecoder(r.Body).Decode(note)
	if err2 != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	res := note.UpdateChild(id)
	res["data"] = id
	u.Respond(w, res)
}

var DeleteChild = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res := models.DeleteChild(id)
	res["data"] = id
	u.Respond(w, res)
}
