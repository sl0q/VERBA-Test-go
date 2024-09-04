package database

import (
	"encoding/json"
	"log"
	"os"

	"VERBA-Test/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Credentials struct {
	Host     string
	User     string
	Password string
	DBname   string
	Port     string
	Sslmode  string
	TimeZone string
}

func ParseCredentials(filePath string) (string, error) {
	credsJson, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Failed to open file [" + filePath + "]")
		return "", err
	}

	var creds Credentials
	json.Unmarshal([]byte(credsJson), &creds)

	dsn := "host=" + creds.Host +
		" user=" + creds.User +
		" password=" + creds.Password +
		" dbname=" + creds.DBname +
		" port=" + creds.Port +
		" sslmode=" + creds.Sslmode +
		" TimeZone=" + creds.TimeZone

	return dsn, nil
}

func ConnectDB(credentialsFilePath string) (*gorm.DB, error) {
	dsn, err := ParseCredentials(credentialsFilePath)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to Database")
		return nil, err
	}

	DB = db
	db.AutoMigrate(&models.Task{})

	return db, nil
}
