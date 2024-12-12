package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB // vai ser um ponteiro para o gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize the SQLite database
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing SQLite database: %v", err)
	}

	return nil
}

func GetSLQite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize the Logger
	logger = NewLogger(p)
	return logger
}
