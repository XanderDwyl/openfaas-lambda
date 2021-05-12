package svc

import (
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3GetFromKeyWithConfig...
func S3GetFromKeyWithConfig(key string, bucket string, s3CFG *s3.S3, decompress bool) ([]byte, error) {
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	item, err := s3CFG.GetObject(params)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(item.Body)
	if decompress {
		rc, err := DecompressBas64(body)
		if err != nil {
			return nil, err
		}

		ba, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}

		return ba, nil
	}

	return body, nil
}
