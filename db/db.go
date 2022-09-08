package db

import (
	"log"

	"github.com/go-redis/redis"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// type Database struct{
// 	DB *sqlx.DB
// }

type Service struct {
	Client *redis.Client
	 Db *sqlx.DB
}


func NewServiceConn(Redis string,dsn string) *Service{
	db,err := sqlx.Connect("postgres",dsn)
	if err != nil {
		log.Fatalln(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: Redis,
		Password: "",
		DB: 0,
	})
	return &Service{
		Client: client,
		Db: db,
	}	

}