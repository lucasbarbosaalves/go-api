package config

import (
	"github.com/lucasbarbosaalves/go-api/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQL() (*gorm.DB, error) {
	logger := GetLogger("mysql")

	// MySQL Configuration
	dsn := "root:1234@tcp(localhost:3306)/gopportunities?charset=utf8mb4&parseTime=True&loc=Local"

	// Create a connection to the MySQL database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("MySQL opening error: %v", err)
		return nil, err
	}

	// Migrate the schema if needed
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("MySQL auto migration error: %v", err)
		return nil, err
	}

	// Return the MySQL DB object
	return db, nil
}
