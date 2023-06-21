package author

import (
	"testing"

	"github.com/russellcxl/go-elastic-iam/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestAuthor(t *testing.T) {
	db.Initialise()
	s := newService()

	// t.Run("create", func(t *testing.T) {
	// 	p := types.Author{
	// 		Name:  "Jim",
	// 		Email: "jim@gmail.com",
	// 	}
	// 	p1, err := s.Save(p)
	// 	assert.Empty(t, err)
	// 	if p1 != nil {
	// 		assert.Equal(t, p.Name, p1.Name)
	// 		assert.Equal(t, p.Email, p1.Email)
	// 	}
	// })

	t.Run("get", func(t *testing.T) {
		_, err := s.Get(1)
		assert.Empty(t, err)
	})

	t.Run("get all", func(t *testing.T) {
		_, err := s.GetAll()
		assert.Empty(t, err)
	})
}
