package db_test

import (
	"testing"

	"github.com/russellcxl/go-elastic-iam/pkg/db"
)

func TestConnect(t *testing.T) {
	db.Connect()
}