package serializer

import "singo/model"

// Video 用户序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(video model.Video) Video {
	return Video{
		ID:        video.ID,
		Title:     video.Title,
		Info:      video.Info,
		CreatedAt: video.CreatedAt.Unix(),
	}
}

// BuildVideoList 序列化视频列表
func BuildVideoList(videos []model.Video) (result []Video) {
	for _, video := range videos {
		result = append(result, BuildVideo(video))
	}
	return result
}
