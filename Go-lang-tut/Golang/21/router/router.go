package router

import (
	"restApi/db"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getAll/platforms", db.GetAllPlateforms).Methods("GET")
	router.HandleFunc("/add/platforms", db.InsertPlateform).Methods("POST")
	router.HandleFunc("/update/platforms/{id}", db.UpdatePlateform).Methods("PUT")
	router.HandleFunc("/delete/platforms/{id}", db.DeletePlateform).Methods("DELETE")
	router.HandleFunc("/deleteAll/platforms", db.DeleteAllPlateforms).Methods("DELETE")
	return router
}
