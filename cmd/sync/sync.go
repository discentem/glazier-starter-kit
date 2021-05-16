package sync

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	return strings.HasPrefix(path, ".git")
}

func configsAndResources(root string, ignore ignorePathFunc) ([]string, error) {
	var names []string
	err := filepath.Walk(root,
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

func Execute(bucketName, accessKey, secretKey, region, root string) error {
	// Retrieve AWS Access Key and Secret Key from env variables
	key := accessKey
	secret := secretKey
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Region:      aws.String(region),
	}
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(s3Config))
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	// value of cnr: [resources/logo.gif stable/config/build.yaml stable/release-id.yaml stable/release-info.yaml version-info.yaml]
	cnr, err := configsAndResources(root, ignorePath) // ignorePath is a function which decides what types of files are NOT considered glazier configs or resources
	if err != nil {
		return err
	}
	// Loop through all the Glazier configs and resource files
	for _, c := range cnr {
		f, err := os.Open(c)
		if err != nil {
			return fmt.Errorf("failed to open file %q, %v", c, err)
		}
		// Upload the config/resource to S3.
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			// S3 does not have folders in the traditional sense. The key represents the entire "path" up to and including the name of the object.
			Key:       &c,
			Body:      f,
			GrantRead: aws.String(`uri="http://acs.amazonaws.com/groups/global/AllUsers"`),
		})
		if err != nil {
			return fmt.Errorf("failed to upload file, %v", err)
		}
		log.Printf("file uploaded to: %s\n", aws.StringValue(&result.Location))
	}
	return nil
}
