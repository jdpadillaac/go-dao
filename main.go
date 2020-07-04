package main

import (
	"fmt"
	"log"

	"./dao/factory"
	"./models"
	"./utilities"
)

func main() {

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	userDao := factory.FactoryDao(config.Engine)
	user := models.User{}

	fmt.Print("Primer nombre:")
	fmt.Scan(&user.FirstName)

	fmt.Print("Primer apellido:")
	fmt.Scan(&user.LastName)

	fmt.Print("Email:")
	fmt.Scan(&user.Email)

	err = userDao.Create(&user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
