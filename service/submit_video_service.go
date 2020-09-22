package service

import (
	"singo/model"
	"singo/serializer"
)

// SubmitVideoService 视频投稿服务
type SubmitVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=255"`
	Info  string `form:"info" json:"info" binding:"max=3000"`
	URL   string `form:"url" json:"url"`
}

// Submit 投稿
func (service *SubmitVideoService) Submit() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
		URL:   service.URL,
	}
	table := model.DB.HasTable(&video)
	if !table {
		model.DB.CreateTable(&video)
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "视频提交失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
