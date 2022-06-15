package internal

/*
Consists all the function to interact to the db
*/
import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

/*
get list of all the cars along its details
*/
func GetCarsData(resp http.ResponseWriter, req *http.Request) {
	carList, err := readAllCarsFromMysql()
	if err != nil {
		returnErrorResponse(http.StatusInternalServerError, "Internal server error", resp)
		return
	}
	resp.WriteHeader(200)
	respJson, _ := json.Marshal(carList)
	resp.Write(respJson)
}

/*
get car detials by car id
*/
func GetCarsDataById(resp http.ResponseWriter, req *http.Request) {
	reqVars := mux.Vars(req)
	reqId := reqVars["id"]

	if isEmpty(reqId) {
		returnErrorResponse(http.StatusBadRequest, "Invalid argument value", resp)
		return
	}
	car, err := readCarsByIdFromMysql(reqId)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("Internal Server Error: " + err.Error()))
		log.Error("Internal server error", err)
		return
	}
	if isEmpty(car.Id) {
		log.Debug("CarId not found ", reqId)
		returnErrorResponse(http.StatusNotFound, "CarId not found", resp)
		return
	}
	resp.WriteHeader(200)
	respJson, _ := json.Marshal(car)
	resp.Write(respJson)
}

//AddNewCar add a new car to Db
func AddNewCar(resp http.ResponseWriter, req *http.Request) {
	var car Cars

	reqData, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqData, &car)
	log.Debug("Car info received: ", car)

	if isEmpty(car.Id) {
		log.Debug("Invalid request data ", car)
		returnErrorResponse(http.StatusBadRequest, "CarId can not be empty", resp)
		return
	}

	err := insertCarIntoMysql(car)
	if err != nil {
		log.Error("Internal server error", err)
		returnErrorResponse(http.StatusInternalServerError, "Internal server error", resp)
		return
	}
	resp.WriteHeader(200)
	respJson, _ := json.Marshal(Response{Code: 200, Message: "Success"})
	resp.Write(respJson)
}
