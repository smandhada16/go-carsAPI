package internal

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func readAllCarsFromMysql() (cars []Cars, err error) {
	res, err := mysqlDb.Query(GET_ALL_CARS_QUERY)
	if err != nil {
		fmt.Println("ERROR: select car details failed. ", err)
		log.Error("ERROR: select car details failed.")
		return cars, err
	}
	for res.Next() {
		car := Cars{}
		res.Scan(&car.Id, &car.Make, &car.Model, &car.Package, &car.Color, &car.Year, &car.Category, &car.Mileage, &car.Price)
		fmt.Println("Car Description: ", car)
		log.Debug("Car Description: ", car)
		cars = append(cars, car)
	}
	return cars, err
}

func readCarsByIdFromMysql(id string) (car Cars, err error) {
	res, err := mysqlDb.Query(GET_CAR_BY_ID_QUERY, id)
	if err != nil {
		fmt.Println("ERROR: select car details failed. ", err)
		log.Error("ERROR: select car details failed. ", err)
		return car, err
	}

	for res.Next() {
		car = Cars{}
		res.Scan(&car.Id, &car.Make, &car.Model, &car.Package, &car.Color, &car.Year, &car.Category, &car.Mileage, &car.Price)
		fmt.Println("Car Description: ", car)
		log.Debug("Car Description: ", car)
		return car, err
	}
	return
}

func insertCarIntoMysql(car Cars) (err error) {
	fmt.Println("Inserting values into Mysql")

	//insert into cars table
	_, err = mysqlDb.Exec(ADD_CAR_QUERY, car.Id, car.Make, car.Model, car.Package, car.Color, car.Year, car.Category, car.Mileage, car.Price)
	if err != nil {
		fmt.Println("ERROR: insert on Car failed. ", err)
		log.Error("ERROR: insert on Car failed. ", err)
		return err
	}

	fmt.Println("Car inserted")
	log.Info("record inserted successfully")
	return
}
