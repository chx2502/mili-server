package service

import (
	"singo/model"
	"singo/serializer"
)

// DeleteVideoService 删除视频的服务
type DeleteVideoService struct {
}

// Delete 删除视频
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	var err error
	err = model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "要删除的视频不存在",
			Error: err.Error(),
		}
	}

	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "删除视频失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 0,
		Msg:  "删除成功",
	}
}
