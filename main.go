package main

import (
	"fmt"
	"log"
	"sqlite-lab/controllers"
	"sqlite-lab/models"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := models.InitialSqliteDatabase(); err != nil {
		log.Panic("initial sqlite error: ", err)
	}
	fmt.Println("sqlite connection success!")
}

func main() {
	r := gin.Default()

	r.POST("/create", controllers.CreateUser)
	r.POST("/update", controllers.UpdateUser)

	r.GET("/get", controllers.GetUserByID)
	r.GET("/list", controllers.GetAllUser)
	r.GET("/delete", controllers.DeleteUser)
	r.GET("/undelete", controllers.UnDeleteUser)

	r.Run(":80")
}
