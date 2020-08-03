package service

import (
	"singo/model"
	"singo/serializer"
)

// ListVideoService 视频列表服务
type ListVideoService struct {
}

// List 创建视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "查询数据库错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideoList(videos),
	}
}
