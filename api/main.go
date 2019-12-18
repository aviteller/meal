package main

import (
	"fmt"
	"net/http"

	auth "./auth"
	"./controllers"
	m "./models"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"},
	})
	m.InitTables()

	//Child routes
	router.HandleFunc("/api/child", controllers.CreateChild).Methods("POST")
	router.HandleFunc("/api/child", controllers.GetChildren).Methods("GET")
	router.HandleFunc("/api/child/{id}", controllers.UpdateChild).Methods("PUT")
	router.HandleFunc("/api/child/{id}", controllers.DeleteChild).Methods("DELETE")
	router.HandleFunc("/api/children", controllers.GetChildrenByUser).Methods("GET")

	// user routes
	router.HandleFunc("/api/user/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/password/{id}", controllers.UpdateUserPassword).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE")

	// meal routes
	router.HandleFunc("/api/meal", controllers.CreateMeal).Methods("POST")
	router.HandleFunc("/api/meal/{id}", controllers.GetMeal).Methods("GET")
	router.HandleFunc("/api/meals", controllers.GetMealsByUser).Methods("GET")
	router.HandleFunc("/api/meal/{id}", controllers.DeleteMeal).Methods("DELETE")

	router.HandleFunc("/api/ingredient/{id}", controllers.DeleteIngredient).Methods("DELETE")
	router.HandleFunc("/api/ingredient", controllers.CreateIngredient).Methods("POST")

	//general
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static/public").HTTPBox()))
	router.Use(auth.JwtAuthentication)

	port := "9000"
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, c.Handler(router))
	if err != nil {
		fmt.Println(err)
	}

}
