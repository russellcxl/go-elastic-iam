package video

import "github.com/russellcxl/go-elastic-iam/pkg/types"

type VideoService interface {
	Save(types.Video) types.Video
	FindAll() []types.Video
}

type videoService struct {
	videos []types.Video
}

func NewService() VideoService {
	return new(videoService)
}

func (s *videoService) Save(v types.Video) types.Video {
	s.videos = append(s.videos, v)
	return v
}

func (s *videoService) FindAll() []types.Video {
	return s.videos
}