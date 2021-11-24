package storage

import (
	"fmt"
	"io/ioutil"
)

type S3Provider struct {

}

func (s3 S3Provider) Upload(request UploadRequest) (string, error){
	// TODO send file to s3
	fmt.Println(request.Key)
	fmt.Println("TODO upload to s3")
	path := "path"+request.Key
	b, _ := ioutil.ReadFile(request.Body.Name()) 
	fmt.Println(string(b))
	return path, nil
}

func NewS3Provider() (*S3Provider) {
	return &S3Provider{}
}