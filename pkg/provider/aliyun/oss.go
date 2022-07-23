package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudstation/pkg/store"
	"github.com/go-playground/validator"
)

var (
	validate = validator.New()
)

type aliyun struct {
	Endpoint  string `validate:"required"`
	AccessKey string `validate:"required"`
	SecretKey string `validate:"required"`
}

func (p *aliyun) UploadFile(bucketName, objectKey, localFile string) error {

	bucket, err := p.GetBucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(objectKey, localFile)
	if err != nil {
		return fmt.Errorf("upload file to bucket: %s error, %s", bucketName, err)
	}

	return nil

}

func (p *aliyun) GetBucket(bucketName string) (*oss.Bucket, error) {
	if bucketName == "" {
		return nil, fmt.Errorf("upload bucket name required")
	}

	// New client
	client, err := oss.New(p.Endpoint, p.AccessKey, p.SecretKey)
	if err != nil {
		return nil, err
	}
	// Get bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func NewUploader(endpoint, accessKey, secretKey string) (store.Uploader, error) {
	aliyun := &aliyun{
		Endpoint:  endpoint,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	if err := aliyun.validate(); err != nil {
		return nil, err
	}
	return aliyun, nil

}

func (p *aliyun) validate() error {
	return validate.Struct(p)
}
