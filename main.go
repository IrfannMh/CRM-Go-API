package main

import (
	"CRM/modules/admin"
	"CRM/modules/customers"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	dsn := "root:root123@tcp(localhost:3306)/mini_project?parseTime=true"
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
		customerRouter.GET("/:page", customerHandler.GetAll)
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
	}

	// 	create customer
	// create admin
	// approve admin
	// get all approve
	// login admin/superadmin
	// delete customer
	// delete admin
	// update admin
	// get all admin by username

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
