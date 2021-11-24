package storage

import "os"

type UploadRequest struct {
	Bucket string
	Key string
	Body os.File
}

type Provider interface {
	Upload(UploadRequest) (string, error)
}