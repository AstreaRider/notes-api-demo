package config_test

import (
	"testing"

	"github.com/AstreaRider/notes-api-demo/config"
	"github.com/joho/godotenv"
)

func TestConfig(t *testing.T) {
	err := godotenv.Load("../.env")
    if err != nil {
        t.Fatalf("Error loading .env file: %v", err)
    }
	t.Run("Load existing database port from .env file", func(t *testing.T) {
		got := config.Config("DB_PORT")
		
		want := "5432"
		if got != want {
			t.Errorf("Config() = %q, want %q", got, want)
		}
	})
}