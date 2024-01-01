package s3

import (
	"context"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

// UploadFile uploads a file to S3 with context
func (c *Client) UploadFile(ctx context.Context, bucketName string, objectName string, filePath string) error {
	_, err := c.client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

// DownloadFile downloads a file from S3 with context
func (c *Client) DownloadFile(ctx context.Context, bucketName string, objectName string, filePath string) error {
	err := c.client.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

// RemoveFile removes a file from S3 with context
func (c *Client) RemoveFile(ctx context.Context, bucketName string, objectName string) error {
	err := c.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

// ListFiles lists all files in a bucket with context
func (c *Client) ListFiles(ctx context.Context, bucketName string) ([]string, error) {
	var files []string

	for object := range c.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{}) {
		if object.Err != nil {
			return nil, object.Err
		}
		files = append(files, object.Key)
	}

	return files, nil
}

// FileExists checks if a file exists with context
func (c *Client) FileExists(ctx context.Context, bucketName string, objectName string) (bool, error) {
	_, err := c.client.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetFileURL returns a URL for a file with context
func (c *Client) GetFileURL(ctx context.Context, bucketName string, objectName string) (string, error) {
	presignedURL, err := c.client.PresignedGetObject(ctx, bucketName, objectName, time.Second*24*60*60, url.Values{})
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
