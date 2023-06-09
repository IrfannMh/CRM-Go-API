package main

import (
	"CRM/modules/admin"
	"CRM/modules/customers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	dbPort := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DBNAME")
	//dsn := "root:root123@tcp(localhost:3306)/crm?parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbPort, dbname)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {

	db, err := initDB()
	if err != nil {
		log.Fatalln("initDB:", err)
	}

	r := gin.Default()
	customerHandler := customers.DefaultRequestHandler(db)

	customerRouter := r.Group("/customers")
	{
		customerRouter.POST("/", customerHandler.Create)
		customerRouter.GET("/:email", customerHandler.GetAll)
		customerRouter.DELETE("/:id", customerHandler.Delete)
	}

	adminHandler := admin.DefaultRequestHandlerAdmin(db)
	adminRouter := r.Group("/admin")
	{
		adminRouter.POST("/", adminHandler.Create)
		adminRouter.GET("/approve", adminHandler.GetAllApprove)
		adminRouter.GET("/:username", adminHandler.GetAdminByUsername)
		adminRouter.PUT("/approve/:id", adminHandler.ApproveByID)
		adminRouter.PUT("/active/:id", adminHandler.ActiveAdmin)
		adminRouter.POST("/login", adminHandler.Login)
		adminRouter.DELETE("/:id", adminHandler.DeleteAdminByID)
	}

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
