package controller

import (
	"backend/middleware"
	"backend/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", routes.Register).Methods("POST")
	router.HandleFunc("/login", routes.Login).Methods("POST")
	router.HandleFunc("/refresh", routes.RefreshToken).Methods("POST")
	
	router.Handle("/getAllChats/{id}", middleware.VerifyToken(http.HandlerFunc(routes.GetAllChats))).Methods("GET")

	router.Handle("/getChat/{id}", middleware.VerifyToken(http.HandlerFunc(routes.GetChat))).Methods("GET")

	router.Handle("/deleteChat/{id}", middleware.VerifyToken(http.HandlerFunc(routes.DeleteChat))).Methods("GET")

	router.Handle("/createChat", middleware.VerifyToken(http.HandlerFunc(routes.CreateChat))).Methods("POST")

	router.Handle("/deleteAllChats/{id}", middleware.VerifyToken(http.HandlerFunc(routes.DeleteAllChats))).Methods("GET")

	router.Handle("/getResponse", middleware.VerifyToken(http.HandlerFunc(routes.GetResponse))).Methods("POST")

	return router
}
