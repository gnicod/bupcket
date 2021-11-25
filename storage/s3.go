package storage

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Provider struct {
	session *session.Session
}

func (s S3Provider) Upload(request UploadRequest) (string, error){
	// TODO send file to s3
	fmt.Println(request.Key)
	fmt.Println("TODO upload to s3")
	path := "path"+request.Key
	b, _ := ioutil.ReadFile(request.Body.Name()) 
	fmt.Println(string(b))
	_, err := s3.New(s.session).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(request.Bucket),
		Key:                aws.String(request.Key),
		ACL:                aws.String("public-read"),
		Body:               &request.Body, // bytes.NewReader(buffer),
		ContentDisposition: aws.String("attachment"),
		// ContentLength:      aws.Int64(int64(len(buffer))),
		// ContentType:        aws.String(http.DetectContentType(buffer)),
		// ServerSideEncryption: aws.String("AES256"),
	})
	return path, err
}

func NewS3Provider() (*S3Provider) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-3")})
	if err != nil {
		log.Fatalf("session.NewSession, err: %v", err)
	}
	return &S3Provider{
		session: sess,
	}
}