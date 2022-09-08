package main

import (
	"github.com/VJ-Vijay77/redis_redigo_pkg/api"
	"github.com/VJ-Vijay77/redis_redigo_pkg/db"
	"github.com/VJ-Vijay77/redis_redigo_pkg/models"
	"github.com/VJ-Vijay77/redis_redigo_pkg/routes"
)

var K *db.Service
const DSN = "postgresql://vijay:12345@localhost:5432/redis?sslmode=disable"
func main() {
	e := routes.EchoService()
	K = db.NewServiceConn("localhost:6379",DSN)
	api.EchoApi(e)
	K.Db.Exec(models.Schema)
	e.Logger.Fatal(e.Start(":8080"))
}
