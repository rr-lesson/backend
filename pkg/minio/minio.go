package minio

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func New() *minio.Client {
	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	minioUser := os.Getenv("MINIO_USER")
	minioPassword := os.Getenv("MINIO_PASSWORD")
	minioUseSSL := os.Getenv("MINIO_USE_SSL") == "true"

	client, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioUser, minioPassword, ""),
		Secure: minioUseSSL,
	})
	if err != nil {
		log.Fatalf("failed to connect minio: %v", err)
	}

	return client
}
