package storage

type UploadRequest struct {
	Bucket string
	Key    string
	Body   []byte
}

type UploadResponse struct {
	Path string `json:"path"`
}

type Provider interface {
	Upload(UploadRequest) (UploadResponse, error)
}
