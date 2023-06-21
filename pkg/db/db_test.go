package db

import (
	"testing"

)

func TestDB(t *testing.T) {
	Initialise()
	migrate()
}
