package main

import (
	"fmt"
	"net/http"

	auth "./auth"
	"./controllers"
	m "./models"
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
	router.HandleFunc("/api/user/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/child/new", controllers.CreateChild).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/children", controllers.GetChildren).Methods("GET")

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
