package config_test

import (
	"os"
	"testing"

	"github.com/AstreaRider/notes-api-demo/config"
	"github.com/joho/godotenv"
)

func TestConfig(t *testing.T) {
	t.Run("Handle missing .env file", func(t *testing.T) {
        if _, err := os.Stat("../.env"); err == nil {
            os.Rename("../.env", "../.env.bak")
            defer os.Rename("../.env.bak", "../.env") 
        }

        
        got := config.Config("DB_PORT")
        want := "" 
        if got != want {
            t.Errorf("Config() when .env is missing = %q, want %q", got, want)
        }
    })

	t.Run("Load existing database port from .env file", func(t *testing.T) {
		err := godotenv.Load("../.env")
		if err != nil {
			t.Fatalf("Error loading .env file: %v", err)
		}

		got := config.Config("DB_PORT")
		want := "5432"
		if got != want {
			t.Errorf("Config() = %q, want %q", got, want)
		}
	})

    t.Run("Handle missing environment variable", func(t *testing.T) {
        got := config.Config("NON_EXISTENT_VAR")
        want := "" 
        if got != want {
            t.Errorf("Config() with non-existent variable = %q, want %q", got, want)
        }
    })
}