package db

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestDB(t *testing.T) {
	godotenv.Load("../../.env")
	Initialise()
	migrate()
}
