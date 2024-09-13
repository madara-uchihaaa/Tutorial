package main

import (
	"backend/controller"
	"net/http"
)

func main() {
	router := controller.Router()
	panic(http.ListenAndServe(":8000", router))
}
