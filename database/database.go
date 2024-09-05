package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
	"open-btm.com/configs"
)

var (
	DBConn *gorm.DB
)

func GormLoggerFile() *os.File {

	gormLogFile, gerr := os.OpenFile("gormblue.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if gerr != nil {
		log.Fatalf("error opening file: %v", gerr)
	}
	return gormLogFile
}

func ReturnSession() (*gorm.DB, error) {

	//  setting up database connection based on DB type

	app_env := configs.AppConfig.Get("DB_TYPE")
	//  This is file to output gorm logger on to
	gormlogger := GormLoggerFile()
	gormFileLogger := log.Logger{}
	gormFileLogger.SetOutput(gormlogger)
	gormFileLogger.Writer()

	gormLogger := log.New(gormFileLogger.Writer(), "\r\n", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile)
	newLogger := logger.New(
		gormLogger, // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			Colorful:                  true,        // Enable color
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			// ParameterizedQueries:      true,        // Don't include params in the SQL log

		},
	)

	var DBSession *gorm.DB

	switch app_env {
	case "postgres":
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  configs.AppConfig.Get("POSTGRES_URI"),
			PreferSimpleProtocol: true, // disables implicit prepared statement usage,

		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   newLogger,
			SkipDefaultTransaction:                   true,
		})
		if err != nil {
			panic(err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("Error during connecting to database: %v\n", err)
			return nil, err
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(5 * time.Second)

		DBSession = db
	case "sqlite":
		//  this is sqlite connection
		db, _ := gorm.Open(sqlite.Open(configs.AppConfig.Get("SQLLITE_URI")), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   newLogger,
			SkipDefaultTransaction:                   true,
		})

		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("Error during connecting to database: %v\n", err)
			return nil, err
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(5 * time.Second)
		DBSession = db
	case "mysql":
		db, _ := gorm.Open(mysql.New(mysql.Config{
			DSN:                       configs.AppConfig.Get("MYSQL_URI"), // data source name
			DefaultStringSize:         256,                                // default size for string fields
			DisableDatetimePrecision:  true,                               // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,                               // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,                               //  when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false,                              // auto configure based on currently MySQL version
		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   newLogger,
			SkipDefaultTransaction:                   true,
		})

		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("Error during connecting to database: %v\n", err)
			return nil, err
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(5 * time.Second)
		DBSession = db
	case "":
		return nil, fmt.Errorf("Error not enough information about connectionis Provided\n")
	default:
		return nil, fmt.Errorf("Error not enough information about connectionis Provided\n")

	}

	DBSession.Use(tracing.NewPlugin())
	return DBSession, nil

}
func ReturnSessionDatabase(dbname string) (*gorm.DB, error) {

	//  setting up database connection based on DB type

	app_env := configs.AppConfig.Get("DB_TYPE")
	//  This is file to output gorm logger on to
	gormlogger := GormLoggerFile()
	gormFileLogger := log.Logger{}
	gormFileLogger.SetOutput(gormlogger)
	gormFileLogger.Writer()

	gormLogger := log.New(gormFileLogger.Writer(), "\r\n", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile)
	newLogger := logger.New(
		gormLogger, // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			Colorful:                  true,        // Enable color
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			// ParameterizedQueries:      true,        // Don't include params in the SQL log

		},
	)

	var DBSession *gorm.DB

	switch app_env {
	case "postgres":
		dsn := fmt.Sprintf("%v dbname=%v sslmode=disable", configs.AppConfig.Get("POSTGRES_URI_NO_DB"), dbname)
		fmt.Println(dsn)
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage,

		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   newLogger,
			SkipDefaultTransaction:                   true,
		})
		if err != nil {
			panic(err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("Error during connecting to database: %v\n", err)
			return nil, err
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(5 * time.Second)

		DBSession = db
	case "sqlite":
		//  this is sqlite connection
		dsn := fmt.Sprintf("%v %v", configs.AppConfig.Get("SQLLITE_URI_NO_DB"), dbname)
		fmt.Println(dsn)
		db, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   newLogger,
			SkipDefaultTransaction:                   true,
		})

		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("Error during connecting to database: %v\n", err)
			return nil, err
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(5 * time.Second)
		DBSession = db
	case "mysql":
		dsn := fmt.Sprintf("%v %v", configs.AppConfig.Get("MYSQL_URI_NO_DB"), dbname)
		db, _ := gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,   // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  //  when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   newLogger,
			SkipDefaultTransaction:                   true,
		})

		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("Error during connecting to database: %v\n", err)
			return nil, err
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(5 * time.Second)
		DBSession = db
	case "":
		return nil, fmt.Errorf("Error Connecting to Provided Sqlite")
	default:
		return nil, fmt.Errorf("Error Connecting to Provided Sqlite")

	}

	DBSession.Use(tracing.NewPlugin())
	return DBSession, nil

}

func CreateDatabase(dbname string) (bool, error) {

	//  setting up database connection based on DB type
	app_env := configs.AppConfig.Get("DB_TYPE")
	switch app_env {
	case "postgres":
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  configs.AppConfig.Get("POSTGRES_URI_NO_DB"),
			PreferSimpleProtocol: true, // disables implicit prepared statement usage,

		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   true,
		})
		if err != nil {
			return false, fmt.Errorf("Error Connecting to Postgres: %v", err)
		}

		db.Exec(fmt.Sprintf("CREATE DATABASE %v", dbname))
		return true, nil
	case "sqlite":
		//  this is sqlite connection
		db, err := gorm.Open(sqlite.Open(fmt.Sprintf(configs.AppConfig.Get("SQLITE_URI_DB_PLACE_HOLDER"), dbname)), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   true,
		})

		if err != nil {
			return false, fmt.Errorf("Error Connecting to Sqlite: %v", err)
		}
		db.Exec(fmt.Sprintf("CREATE DATABASE %v", dbname))
		return true, nil
	case "mysql":
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       configs.AppConfig.Get("MYSQL_URI"), // data source name
			DefaultStringSize:         256,                                // default size for string fields
			DisableDatetimePrecision:  true,                               // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,                               // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,                               //  when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false,                              // auto configure based on currently MySQL version
		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   true,
		})
		if err != nil {
			return false, fmt.Errorf("Error Connecting to Mysql: %v", err)
		}

		db.Exec(fmt.Sprintf("CREATE DATABASE %v", dbname))
		return true, nil
	case "":
		return false, fmt.Errorf("No Database Name or DB type(on env file) Provided")
	default:
		return true, fmt.Errorf("No Data base type or Name provided")
	}
}
