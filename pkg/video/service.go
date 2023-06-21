package video

import (
	"github.com/russellcxl/go-elastic-iam/pkg/db"
	"github.com/russellcxl/go-elastic-iam/pkg/types"
	"gorm.io/gorm"
)

type VideoService interface {
	Save(types.Video) types.Video
	FindAll() []types.Video
}

type videoService struct {
	db *gorm.DB
	videos []types.Video
}

func NewService() VideoService {
	return &videoService{
		db: db.DB,
	}
}

func (s *videoService) Save(v types.Video) types.Video {
	s.videos = append(s.videos, v)
	return v
}

func (s *videoService) FindAll() []types.Video {
	return s.videos
}