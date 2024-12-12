package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB // vai ser um ponteiro para o gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	// Initialize the Logger
	logger = NewLogger(p)
	return logger
}
