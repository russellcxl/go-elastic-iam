package db

import (
	"github.com/russellcxl/go-elastic-iam/pkg/types"
)

func migrate() {
	if err := DB.AutoMigrate(
		&types.Video{},
		&types.Author{},
	); err != nil {
		panic("failed to migrate: " + err.Error())
	}
}
