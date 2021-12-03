package main

import "github.com/gorilla/mux"

func setRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/list/", list).Methods("GET")
	router.HandleFunc("/retrieve/{UUID}", retrieve).Methods("GET")
	router.HandleFunc("/add/", add).Methods("POST")
	router.HandleFunc("/delete/{UUID}", deleteData).Methods("DELETE")
	router.HandleFunc("/update/", update).Methods("PUT")
	return router
}
