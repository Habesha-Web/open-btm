package models

import (
	"fmt"
	"log"

	"open-btm.com/configs"
	"open-btm.com/database"
)

func InitDatabase() {
	app_env := configs.AppConfig.Get("APP_ENV")
	if app_env != "test" {
		configs.NewEnvFile("./configs")
	}

	database, err := database.ReturnSession()
	fmt.Println("Connection Opened to Database")
	if err == nil {
		if err := database.AutoMigrate(
			&Project{},
			&ProjectUsers{},
		); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Database Migrated")
	} else {
		panic(err)
	}
}

func MigrateToPojectDatabase(dbname string) {
	configs.NewEnvFile("./configs")
	database, err := database.ReturnSessionDatabase(dbname)
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
		fmt.Printf("Database Migrated to : %v\n", dbname)
	} else {
		panic(err)
	}
}

func CleanDatabase() {
	configs.NewEnvFile("./configs")
	mdb, err := database.ReturnSession()
	var projects []Project

	mdb.Model(&Project{}).Find(&projects)

	if err == nil {
		fmt.Println("Connection Opened to Database")
		for _, proj := range projects {
			db, err := database.ReturnSessionDatabase(proj.DatabaseName)
			if err != nil {
				db.Migrator().DropTable(
					&Sprint{},
					&Requirement{},
					&Test{},
					&Testset{},
					&TestTestset{},
					&Issue{},
				)
			}

		}

		fmt.Println("Dropping Models if Exist")
		mdb.Migrator().DropTable(
			&Project{},
			&ProjectUsers{},
		)
		fmt.Println("Database Cleaned")
	} else {
		panic(err)
	}
}
