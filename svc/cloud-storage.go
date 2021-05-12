package svc

import (
	"context"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
)

// GCSGetFromKeyWithConfig...
func GCSGetFromKeyWithConfig(key string, bucket string, decompress bool) ([]byte, error) {

	// Make sure to set GOOGLE_APPLICATION_CREDENTIALS env
	ctx := context.Background()
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client : %v", err)
		return nil, err
	}

	b := storageClient.Bucket(bucket)
	obj := b.Object(key).ReadCompressed(true) // see https://developer.bestbuy.com/apis
	rc, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}

	defer rc.Close()

	body, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

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
