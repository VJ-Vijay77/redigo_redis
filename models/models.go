package models


type Shop struct{
	ShopName string `json:"name"`
	Owner string 	`json:"owner"`
	Item string	 	`json:"item"`
	Price int64 	`json:"price"`
	Pass string		`json:"pass"`
}


var Schema = `CREATE TABLE shop(
	id serial primary key,
	shopname text,
	owner text,
	item text,
	price integer,
	pass text);`