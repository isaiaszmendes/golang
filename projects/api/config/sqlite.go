package config

import (
	"os"

	"github.com/isaiaszmendes/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	// Initialize the SQLite database

	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("SQLite database file does not exist, creating it ...")
		// Create the database file
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating the SQLite database file: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating the SQLite database file: %v", err)
			return nil, err
		}
		err = file.Close()
		if err != nil {
			logger.Errorf("Error closing the SQLite database file: %v", err)
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing the SQLite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating the SQLite database: %v", err)
		return nil, err
	}

	logger.Info("SQLite database initialized")

	// Return the DB
	return db, nil
}
