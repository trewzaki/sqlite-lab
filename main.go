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
}

func main() {
	fmt.Println("sqlite connection success!")

	// userModel := models.User{}

	/* Create */
	// userModel.FirstName = "Warakorn"
	// userModel.LastName = "Meakpattanapinyo"
	// userModel.Create()

	/* Get All */
	// usersResponse, getAllError := userModel.GetAll()
	// if getAllError != nil {
	// 	log.Panic("get all user error: ", getAllError)
	// }
	// fmt.Println("usersResponse: ", usersResponse)

	fmt.Println("done!")
}
