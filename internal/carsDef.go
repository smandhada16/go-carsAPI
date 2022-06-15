package internal

type Cars struct {
	Id       string `json:"id"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Package  string `json:"package"`
	Color    string `json:"color"`
	Year     string `json:"year"`
	Category string `json:"category"`
	Mileage  int    `json:"mileage"`
	Price    int    `json:"price"`
}

const (
	CARS_CREATE_TABLE_QUERY = "create table IF NOT EXISTS Cars (id varchar(250) primary key , make varchar(250), model varchar(250), package varchar(250), color varchar(250),year varchar(250),category varchar(250),mileage int,price int)"
	GET_CAR_BY_ID_QUERY     = "select * from cars where id = ?"
	GET_ALL_CARS_QUERY      = "select * from cars"
	ADD_CAR_QUERY           = "Insert into cars(id, make, model, package, color, year, category, mileage, price) values(?, ?, ?, ?, ?, ?, ?, ?, ?)"
)

type Response struct {
	Code    int    `json:"id"`
	Message string `json:"msg"`
}
