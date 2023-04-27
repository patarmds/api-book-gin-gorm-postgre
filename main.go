package main

import (
	"rest-api-books-gin-gorm/routers"
	"rest-api-books-gin-gorm/database"
)

func main(){
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}