package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nehonar/twitteringo/middlew"
	"github.com/nehonar/twitteringo/routers"
	"github.com/rs/cors"
)

/*
HandlersRoute create routes and cors get controll and gift permision allow all
*/
func HandelersRoute() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", middlew.CheckInMongoDB(routers.SignUpUser)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckInMongoDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
