package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
)

func ConnectToDB() (*gorm.DB, error) {
	var (
		err         error
		maxTries    = 10
		actualTries int
		ticker      = time.NewTicker(time.Second)
		dsn         = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_EXPOSED_PORT"))
	)

	//nolint:gosimple
	for {
		select {
		case <-ticker.C:
			dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			actualTries++
			if err != nil {
				if actualTries == maxTries {
					err = fmt.Errorf("failed to connect to db after %d tries: %v", actualTries, err)
					return nil, err
				}

				continue
			}

			log.Print("db connection created")

			return dbConn, nil
		}
	}
}

func MigrateDB() error {
	if err := dbConn.AutoMigrate(&User{}); err != nil {
		err = fmt.Errorf("failed to run db migrations: %v", err)
		return err
	}

	log.Print("db migrations executed")

	return nil
}
