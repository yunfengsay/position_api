package main

import (
	"./db"
	_ "./models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"position_api/apis"
	"position_api/models"
)

func getDBConnection() *gorm.DB {
	db, err := db.Connect("localhost", "position", "yunfeng_db", "pssword")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Location{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Like{})

	fmt.Println("Connect ok")
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate()
	return db
}

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", apis.IndexApi)
	//
	//router.POST("/person", AddPersonApi)
	//
	//router.GET("/persons", GetPersonsApi)
	//
	//router.GET("/person/:id", GetPersonApi)
	//
	//router.PUT("/person/:id", ModPersonApi)
	//
	//router.DELETE("/person/:id", DelPersonApi)

	return router
}

func main() {
	db := getDBConnection()

	defer db.Close()
	router := initRouter()
	router.Run(":8001")

}
