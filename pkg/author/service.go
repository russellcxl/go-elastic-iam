package author

import (
	"github.com/russellcxl/go-elastic-iam/pkg/db"
	"github.com/russellcxl/go-elastic-iam/pkg/types"
	"gorm.io/gorm"
)

type AuthorService interface {
	Save(types.Author) (*types.Author, error)
	Get(id uint) (*types.Author, error)
	GetAll() ([]types.Author, error)
}

type authorService struct {
	db *gorm.DB
}

func newService() AuthorService {
	return &authorService{
		db: db.DB,
	}
}

func (s *authorService) Save(p types.Author) (*types.Author, error) {
	res := s.db.Save(&p)
	if res.Error != nil {
		return nil, res.Error
	}
	return &p, nil
}

func (s *authorService) Get(id uint) (*types.Author, error) {
	var p types.Author
	res := s.db.First(&p)
	if res.Error != nil {
		return nil, res.Error
	}
	return &p, nil

}

func (s *authorService) GetAll() ([]types.Author, error) {
	var p []types.Author
	res := s.db.Find(&p)
	if res.Error != nil {
		return nil, res.Error
	}
	return p, nil
}

