package models

import (
	"fmt"
	"log"

	"open-btm.com/configs"
	"open-btm.com/database"
)

func InitDatabase() {
	configs.NewEnvFile("./configs")
	database, err  := database.ReturnSession()
	fmt.Println("Connection Opened to Database")
	if err == nil {
		if err := database.AutoMigrate(
			&Sprint{},
			&Requirement{},
			&Test{},
			&Testset{},
			&TestTestset{},
			&Issue{},
		); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Database Migrated")
	} else {
		panic(err)
	}
}

func CleanDatabase() {
	configs.NewEnvFile("./configs")
	database, err := database.ReturnSession()
	if err == nil {
		fmt.Println("Connection Opened to Database")
		fmt.Println("Dropping Models if Exist")
		database.Migrator().DropTable(

			&Sprint{},

			&Requirement{},

			&Test{},

			&Testset{},

			&TestTestset{},

			&Issue{},

		)

		fmt.Println("Database Cleaned")
	} else {
		panic(err)
	}
}
