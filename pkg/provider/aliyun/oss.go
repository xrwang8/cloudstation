package aliyun

import "github.com/cloudstation/pkg/store"

type aliyn struct{}

func (a aliyn) UploadFile(bucketName, objectKey, localFile string) error {
	//TODO implement me
	panic("implement me")
}

func NewUploader() store.Uploader {
	return &aliyn{}

}
