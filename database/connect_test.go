package database_test

import (
	"os"
	"testing"

	"github.com/AstreaRider/notes-api-demo/database"
	"github.com/joho/godotenv"
)

func setupDatabaseTest() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "test_db")
	os.Setenv("DB_USER", "test")
	os.Setenv("DB_PASSWORD", "test")
	os.Setenv("DB_PORT", "5432")
}

func TestConnectDB(t *testing.T ) {
	err := godotenv.Load("../.env")
		if err != nil {
			t.Fatalf("Error loading .env file: %v", err)
		}

	database.ConnectDB()
	if database.DB == nil {
        t.Fatal("Expected DB to be initialized, but it is nil")
    }

    sqlDB, err := database.DB.DB()
    if err != nil {
        t.Fatalf("Failed to get sql.DB: %v", err)
    }

    err = sqlDB.Ping()
    if err != nil {
        t.Fatalf("Failed to ping the database: %v", err)
    }
}