package config

import (
	"database/sql"
	"fmt"
	"log"

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

type DatabaseConnection struct {
	Conn    *sql.DB
	DB      *gorm.DB
	Migrate *migrate.Migrate
}

func StartConnectionDB() (*DatabaseConnection, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		Environment.DB_HOST, Environment.DB_USER, Environment.DB_PASSWORD, Environment.DB_NAME, Environment.DB_PORT,
	)

	db, err := gorm.Open(driverGorm.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("failed to connect to the Database. \n", err.Error())
		return nil, err
	}

	con, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get DB object. \n", err.Error())
		return nil, err
	}

	if Environment.ENV == DEV {
		db.Logger = logger.Default.LogMode(logger.Info)
	} else {
		db.Logger = logger.Default.LogMode(logger.Silent)
	}

	driver, err := postgres.WithInstance(con, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Postgres migrate from current connection. \n", err.Error())
		return nil, err
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
		return nil, err
	}

	return &DatabaseConnection{
		Conn:    con,
		DB:      db,
		Migrate: m,
	}, nil
}

func ConnectDB() {
	connection, err := StartConnectionDB()
	if err != nil {
		log.Fatal("Failed to connect to the Database. \n", err.Error())
	}

	Conn = connection.Conn
	DB = connection.DB

	if !Environment.RUN_MIGRATIONS {
		return
	}

	err = connection.Migrate.Up()
	if err != migrate.ErrNoChange {
		log.Fatal("failed to run all migrations UP: ", err.Error())
	}
}
