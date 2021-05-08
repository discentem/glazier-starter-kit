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

func isFile(path string) (bool, error) {
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

type ignorePathFunc func(path string) bool

func ignorePath(path string) bool {
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

func glazierConfigsAndResources(root string, ignore ignorePathFunc) ([]string, error) {
	var names []string
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if ignore(path) {
				return nil
			}
			isfl, err := isFile(path)
			if err != nil {
				return err
			}
			if isfl {
				names = append(names, path)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return names, nil
}

func main() {
	// Retrieve AWS Access Key and Secret Key from env variables
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
	// Ex. value of configs [resources/logo.gif stable/config/build.yaml stable/release-id.yaml stable/release-info.yaml version-info.yaml]
	configsAndResources, err := glazierConfigsAndResources(".", ignorePath) // ignorePath is a function which decides what types of files are NOT considered glazier configs or resources
	if err != nil {
		log.Fatal(err)
	}
	// Loop through all the Glazier configs and resource files
	for _, c := range configsAndResources {
		f, err := os.Open(c)
		if err != nil {
			log.Fatalf("failed to open file %q, %v", c, err)
		}
		// Upload the config/resource to S3.
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(os.Getenv("BUCKET_NAME")),
			// S3 does not have folders in the traditional sense. The key represents the entire "path" up to and including the name of the object.
			// glazierConfigsAndResources effectively converts file system path into keys, so
			Key:  &c,
			Body: f,
		})
		if err != nil {
			log.Fatalf("failed to upload file, %v", err)
		}
		log.Printf("file uploaded to: %s\n", aws.StringValue(&result.Location))
	}

}
