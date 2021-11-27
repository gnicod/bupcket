package storage

import "os"

type UploadRequest struct {
	Bucket string
	Key    string
	Body   os.File
}

type UploadResponse struct {
	Path string `json:"path"`
}

type Provider interface {
	Upload(UploadRequest) (UploadResponse, error)
}
