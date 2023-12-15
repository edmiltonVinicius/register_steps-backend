package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/edmiltonVinicius/register-steps/api/utils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	driverGorm "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *sql.DB
var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		Environment.DB_HOST, Environment.DB_USER, Environment.DB_PASSWORD, Environment.DB_NAME, Environment.DB_PORT,
	)

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

	Conn = db
	DB = con

	if Environment.ENV == DEV {
		DB.Logger = logger.Default.LogMode(logger.Info)
	} else {
		DB.Logger = logger.Default.LogMode(logger.Silent)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Postgres magrate from current connection. \n", err.Error())
		os.Exit(1)
	}

	var migrations string

	if Environment.ENV == TEST {
		root := utils.GetRootPath()
		migrations = "file://" + root + "/database/migrations"
	} else {
		migrations = "file://./database/migrations"
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrations,
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal("Failed to load settings migration: ", err.Error())
		os.Exit(1)
	}

	if !Environment.RUN_MIGRATIONS {
		return
	}

	err = m.Up()
	if err != migrate.ErrNoChange {
		log.Fatal("Failed to run migration: ", err.Error())
		os.Exit(1)
	}
}
