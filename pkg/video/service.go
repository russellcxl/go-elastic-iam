package video

import (
	"github.com/russellcxl/go-elastic-iam/pkg/db"
	"github.com/russellcxl/go-elastic-iam/pkg/types"
	"gorm.io/gorm"
)

type VideoService interface {
	Save(types.Video) (*types.Video, error)
	GetAll() ([]types.Video, error)
}

type videoService struct {
	db *gorm.DB
}

func NewService() VideoService {
	return &videoService{
		db: db.DB,
	}
}

func (s *videoService) Save(v types.Video) (*types.Video, error) {
	res := s.db.Save(&v)
	if res.Error != nil {
		return nil, res.Error
	}
	return &v, nil
}

func (s *videoService) GetAll() ([]types.Video, error) {
	var v []types.Video
	res := s.db.Preload("Author").Find(&v) // https://gorm.io/docs/preload.html
	if res.Error != nil {
		return nil, res.Error
	}
	return v, nil
}