package adapter

import (
	"context"
	"errors"
)

type ObjectStorage interface {
	Upload(ctx context.Context, bucket, key string, data []byte) error
	Download(ctx context.Context, bucket, key string) ([]byte, error)
	Delete(ctx context.Context, bucket, key string) error
}

type GCSUploadRequest struct {
	Bucket string
	Object string
	Data   []byte
}

type GCSGetRequest struct {
	Bucket string
	Object string
}

type GCSDeleteRequest struct {
	Bucket string
	Object string
}

type GCSClient struct {
	objects map[string][]byte
}

func NewGCSClient() *GCSClient {
	return &GCSClient{
		objects: make(map[string][]byte),
	}
}

func (c *GCSClient) UploadObject(ctx context.Context, req GCSUploadRequest) error {
	if req.Bucket == "" || req.Object == "" {
		return errors.New("bucket and object are required")
	}
	c.objects[req.Bucket+"/"+req.Object] = req.Data
	return nil
}

func (c *GCSClient) GetObject(ctx context.Context, req GCSGetRequest) ([]byte, error) {
	data, ok := c.objects[req.Bucket+"/"+req.Object]
	if !ok {
		return nil, errors.New("object not found")
	}
	return data, nil
}

func (c *GCSClient) DeleteObject(ctx context.Context, req GCSDeleteRequest) error {
	key := req.Bucket + "/" + req.Object
	if _, ok := c.objects[key]; !ok {
		return errors.New("object not found")
	}
	delete(c.objects, key)
	return nil
}

type GCSStorageAdapter struct {
	client *GCSClient
}

func NewGCSStorageAdapter(client *GCSClient) *GCSStorageAdapter {
	return &GCSStorageAdapter{client: client}
}

func (a *GCSStorageAdapter) Upload(ctx context.Context, bucket, key string, data []byte) error {
	return a.client.UploadObject(ctx, GCSUploadRequest{
		Bucket: bucket,
		Object: key,
		Data:   data,
	})
}

func (a *GCSStorageAdapter) Download(ctx context.Context, bucket, key string) ([]byte, error) {
	return a.client.GetObject(ctx, GCSGetRequest{
		Bucket: bucket,
		Object: key,
	})
}

func (a *GCSStorageAdapter) Delete(ctx context.Context, bucket, key string) error {
	return a.client.DeleteObject(ctx, GCSDeleteRequest{
		Bucket: bucket,
		Object: key,
	})
}

var _ ObjectStorage = (*GCSStorageAdapter)(nil)
