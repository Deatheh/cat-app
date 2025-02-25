package repository

import (
	"context"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ConfigMinio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

func MinioConnection(cfg ConfigMinio) (*minio.Client, error) {
	endpoint := "127.0.0.1:9000" // адрес с3
	accessKeyID := "minio"
	// логни
	secretAccessKey := "minio124" // пароль
	useSSL := false
	// Инициализация клиента минио
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return minioClient, nil
}

func CreateBucket(minioS3 *minio.Client, bucketName string) {
	location := "us-east-1"
	err := minioS3.MakeBucket(context.Background(), bucketName,
		minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Проверка, существует ли уже бакет с таким именем
		exists, errBucketExists :=
			minioS3.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			log.Printf("bucket %s already exists\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("bucket %s successfully created\n", bucketName)
	}
}
func RFPutObject(minioClient *minio.Client, bucketName, fileName, filePath, contentType string) error {
	_, err := minioClient.FPutObject(context.Background(),
		bucketName, fileName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}
	return nil
}
func RPresignedGetObject(minioClient *minio.Client, bucketName, filePath string, expiration time.Duration) (*url.URL, error) {
	object, err :=
		minioClient.PresignedGetObject(context.Background(), bucketName,
			filePath, expiration, nil)
	if err != nil {
		return nil, err
	}
	return object, nil
}
