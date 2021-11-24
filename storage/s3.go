package storage

import "fmt"

type S3Provider struct {

}

func (s3 S3Provider) Upload(request UploadRequest) {
	fmt.Println("TODO upload to s3")
}

func NewS3Provider() (*S3Provider) {
	return &S3Provider{}
}