package main

import (
	"CarsApp/internal"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func StartApiServer() {
	mx := mux.NewRouter()
	siteserver := http.FileServer(http.Dir("./CarUI"))
	// mx.Handle("/", siteserver)

	mx.Handle("/", siteserver)
	mx.PathPrefix("/Content/").Handler(siteserver)
	mx.PathPrefix("/Scripts/").Handler(siteserver)

	mx.HandleFunc("/cars", internal.GetCarsData).Methods(http.MethodGet)
	mx.HandleFunc("/cars/{id}", internal.GetCarsDataById).Methods(http.MethodGet)
	mx.HandleFunc("/cars/{id}", internal.AddNewCar).Methods(http.MethodPost)

	fmt.Println("Serving on 8000")
	log.Debug("Serving on 8000")
	http.ListenAndServe(":8000", mx)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			return
		}
	}()

	//init db server
	internal.InitMysqlConnection()

	//Init api server
	StartApiServer()
}
