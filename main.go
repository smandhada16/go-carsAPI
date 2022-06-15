package main

import (
	"CarsApp/internal"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

/*
refer class 7 in taining folder

Structure -> all exported firelds and json tags done
interface for functionality like insert, readALl, readById -> optional
implement interface on struct -> optional
file 1 -> definitions like struct, db query done
file 2 -> db connectivity done
		Code copy paster from abover eg clss 7
file 3 -> handler
		refer class 7 -> 3 handler
		getAll -> trigger getALl finction n return
		GetById -> readId and return
		insert -> read data and insert
file 4 -> main
*/
func StartApiServer() {
	mx := mux.NewRouter()

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
