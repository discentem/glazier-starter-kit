package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type ignoreFileFunc func(path string) bool

func ignoreFile(path string) bool {
	switch path {
	case "sync.go":
		return true
	case "go.mod":
		return true
	case "go.sum":
		return true
	default:
		if path[0:1] == "." {
			return true
		}
		return false
	}
}

func fileNames(root string, ignore ignoreFileFunc) ([]string, error) {
	var names []string
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if ignore(path) {
				return nil
			}
			names = append(names, path)
			return err
		})
	if err != nil {
		return nil, err
	}
	return names, nil
}

func IsFile(path string) (bool, error) {
	if path == "" {
		return false, fmt.Errorf("IsFile: received empty string to test")
	}
	p, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return !p.IsDir(), nil
}

func main() {
	key := os.Getenv("ACCESS_KEY")
	secret := os.Getenv("SECRET_KEY")

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Region:      aws.String("us-east-1"),
	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(s3Config))
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	configs, err := fileNames(".", ignoreFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(configs)

	for _, c := range configs {
		isFile, err := IsFile(c)
		if err != nil {
			log.Fatal(err)
		}
		if !isFile {
			continue
		}

		f, err := os.Open(c)
		if err != nil {
			log.Fatalf("failed to open file %q, %v", c, err)
		}

		// Upload the file to S3.
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(os.Getenv("BUCKET_NAME")),
			Key:    &c,
			Body:   f,
		})
		if err != nil {
			log.Fatalf("failed to upload file, %v", err)
		}
		fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	}

}
