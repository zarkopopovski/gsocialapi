package main

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

type DBConnection struct {
	db    *sql.DB
	cache *redis.Client
}

func OpenConnectionSession() (dbConnection *DBConnection) {
	dbConnection = new(DBConnection)
	dbConnection.createNewDBConnection()
	dbConnection.createNewCacheConnection()

	return
}

func (dbConnection *DBConnection) createNewDBConnection() (err error) {
	connection, err := sql.Open("mysql", "root@/application_db?charset=utf8")
	if err != nil {
		panic(err)
	}

	fmt.Println("MySQL Connection is Active")
	dbConnection.db = connection

	return
}

func (dbConnection *DBConnection) createNewCacheConnection() (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       10, // use default DB
	})

	fmt.Println("Redis Connection is Active")
	dbConnection.cache = client

	return
}
