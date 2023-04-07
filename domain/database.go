package domain

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	driverGorm "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", Environment.DB_HOST, Environment.DB_USER, Environment.DB_PASSWORD, Environment.DB_NAME, Environment.DB_PORT)

	con, err := gorm.Open(driverGorm.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to the Database. \n", err.Error())
		os.Exit(1)
	}

	db, err := con.DB()
	if err != nil {
		log.Fatal("Failed to get DB object. \n", err.Error())
		os.Exit(1)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Postgres magrate from current connection. \n", err.Error())
		os.Exit(1)
	}

    m, err := migrate.NewWithDatabaseInstance(
        "file://./database/migrations",
        "postgres", 
		driver,
	)

	if err != nil {
		log.Fatal("Failed to load settings migration. \n", err.Error())
		os.Exit(1)
	}
		
	m.Up()
	DB = con
	DB.Logger = logger.Default.LogMode(logger.Info)
}