package video

import (
	"testing"

	"github.com/russellcxl/go-elastic-iam/pkg/db"
)

func TestVideo(t *testing.T) {
	db.Initialise()
	s := NewService()

	t.Run("create", func(t *testing.T) {
		
	})

	t.Run("get", func(t *testing.T) {
		
	})
}