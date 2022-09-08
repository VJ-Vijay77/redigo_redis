package controllers

import (
	"log"

	"github.com/VJ-Vijay77/redis_redigo_pkg/config"
	"github.com/VJ-Vijay77/redis_redigo_pkg/db"
	"github.com/VJ-Vijay77/redis_redigo_pkg/models"
	"github.com/labstack/echo/v4"
)

var DBS = db.NewServiceConn(config.REDIS,config.DSN)

func Home(c echo.Context) error {
	return c.JSON(200, "Welcome to Home Page")
}

func AddUser(c echo.Context) error {
	var data models.Shop

	if err := c.Bind(&data); err != nil {
		log.Fatalln(err)
	}

	query := `INSERT INTO shop (shopname,owner,item,price,pass)VALUES($1,$2,$3,$4,$5)`
	_, err := DBS.Db.Exec(query, data.ShopName, data.Owner, data.Item, data.Price, data.Pass)
	if err != nil {
		return c.JSON(400,"cannot perform the task")
	}
	
	err = DBS.Client.Set(data.Owner,data.Pass,0).Err()
	if err != nil{
		log.Println(err)
	}



	return c.JSON(200,"Data entered successfull...")

}

func GetPass(c echo.Context) error {
	name := c.Param("name")
	val,err := DBS.Client.Get(name).Result()
	if err != nil {
		log.Println(err)
	}
	return c.JSON(200,val)
}