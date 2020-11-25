package main

import (
	"fmt"
	"go-db-learn/infrastructure/database"
	// "go-db-learn/domain/entity"
)



func main() {

	services, err := database.NewRepository("postgres", "postgres", "5432", "localhost", "config_db")

	if err != nil{
		panic(err)
	}
	defer services.Close()

	services.Automigrate()

	// newCountry := entity.Country{
	// 	Name: "Kenya",
	// 	Iso: "KE",
	// 	Numcode: "404",
	// 	Phonecode: "254",
	// 	Currency: "KES",
	// 	Status: "Active",
	// }


	fmt.Println("Am dealing with databases" + services.Name )

}
