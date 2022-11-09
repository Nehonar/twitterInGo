package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nehonar/twitteringo/middleware"
	"github.com/nehonar/twitteringo/routers"
	"github.com/rs/cors"
)

/*
HandlersRoute create routes and cors get controll and gift permision allow all
*/
func HandelersRoute() {
	router := mux.NewRouter()

	router.HandleFunc("/sing-up", middleware.CheckInMongoDB(routers.SignUpUser)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
