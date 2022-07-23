package store

type Uploader interface {
	UploadFile(bucketName, objectKey, localFile string) error
}
