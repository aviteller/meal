package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
)

//MEALS

var CreateMeal = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	m := &models.Meal{}
	err := json.NewDecoder(r.Body).Decode(m)

	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	m.UserID = int(user)
	res := m.Create()
	u.Respond(w, res)
}

var GetMeal = func(w http.ResponseWriter, r *http.Request) {
	// user := r.Context().Value("user").(uint)
	params := mux.Vars(r)
	id := params["id"]
	data := models.GetMeal(id)

	res := u.Message(true, "success")

	res["data"] = data

	u.Respond(w, res)
}

var GetMealsByUser = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	data := models.GetMealsByUser(int(user))

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)
}

var DeleteMeal = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res := models.DeleteMeal(id)
	res["data"] = id
	u.Respond(w, res)
}

//Ingredient
var CreateIngredient = func(w http.ResponseWriter, r *http.Request) {
	i := &models.Ingredient{}

	err := json.NewDecoder(r.Body).Decode(i)

	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	res := i.Create()
	u.Respond(w, res)
}

var DeleteIngredient = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res := models.DeleteIngredient(id)
	res["data"] = id
	u.Respond(w, res)
}
