package service

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"singo/serializer"
)

// UploadTokenService 获取上传 OSS 的 token
type UploadTokenService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 创建 token
func (service *UploadTokenService) Post() serializer.Response {
	endpoint := os.Getenv("OSS_END_POINT")
	accessKeyId := os.Getenv("OSS_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("OSS_ACCESS_KEY_SECRET")
	bucketName := os.Getenv("OSS_BUCKET")

	// 创建 OSSClient 实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "OSS client 创建错误",
			Error: err.Error(),
		}
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "OSS bucket 获取错误",
			Error: err.Error(),
		}
	}

	ext := filepath.Ext(service.Filename)
	// 带可选参数的签名直传，options 设置文件元信息
	// 参考：https://help.aliyun.com/document_detail/88638.html?spm=a2c4g.11186623.2.9.68cb500bYY85ZN#title-a4v-fqx-mpe
	options := []oss.Option{
		oss.ContentType("image/png"),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext
	// 签名直传
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "OSS 签名直传错误",
			Error: err.Error(),
		}
	}

	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "OSS 签名链接获取错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}