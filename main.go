package main

import (
	"fmt"
	"log"
	"sqlite-lab/models"
)

func init() {
	if err := models.InitialSqliteDatabase(); err != nil {
		log.Panic("initial sqlite error: ", err)
	}
	fmt.Println("sqlite connection success!")
}

func main() {
	// userModel := models.User{}

	/* Create */
	// userModel.FirstName = "Tananut"
	// userModel.LastName = "Panyagosa"
	// userModel.Create()

	/* Get All */
	// usersResponse, getAllError := userModel.GetAll()
	// if getAllError != nil {
	// 	log.Panic("get all user error: ", getAllError)
	// }
	// fmt.Println("usersResponse: ", usersResponse)

	fmt.Println("done!")
}
