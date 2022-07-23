package aliyun

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	bucketName = "cloud-station-demo"
	objectKey  = "oss.go"
	localFile  = "oss.go"

	endPoint  = "oss-cn-hangzhou.aliyuncs.com"
	accessKey = "LTAI5tDGiHxToUz6qB8oxybJ"
	secretKey = "2ADkhVp64iVMTljQbigOYpVuC9AXTV"
)

func TestUploadFile(t *testing.T) {

	should := assert.New(t)
	uploader, err := NewUploader(endPoint, accessKey, secretKey)

	if should.NoError(err) {
		err = uploader.UploadFile(bucketName, objectKey, localFile)
		should.NoError(err)

	}

}
