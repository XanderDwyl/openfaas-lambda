package svc

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"


	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
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
	obj := b.Object(key).ReadCompressed(true)
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

// GCSUpload ...
func GCSUpload( item interface{}, ctype, bucket string, fileName string, compress bool) error {
	var data string

	// Make sure to set GOOGLE_APPLICATION_CREDENTIALS env
	ctx := context.Background()
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client : %v", err)
		return err
	}

	b := storageClient.Bucket(bucket)
	if _, err := b.Attrs(ctx); err == storage.ErrBucketNotExist {
		LogIt("The bucket does not exist")
		return err
	}

	wc := b.Object(fileName).NewWriter(ctx)

	defer wc.Close()

	if compress {
		data, err = Base64Compress(item)
		if err != nil {
			LogIt("Error in Base64Compress")
			return err
		}
	} else {
		var b []byte

		b, err = json.Marshal(item)
		if err != nil {
			LogIt("Error in JSON Marshal")
			return err
		}

		data = string(b)
		ctype = "application/json"
	}

	wc.ContentType = ctype
	wc.RetentionExpirationTime = time.Now().Add(time.Second * 60) // 1-minute
	if _, err := wc.Write([]byte(data)); err != nil {
		LogIt("Error writing content in bucket")
		return err
	}

	return err
}

func GetWithSignedUrl(sakeyFile, bucket, filename string, expiredInMin int) (url string, err error){

	saKey, err := ioutil.ReadFile(sakeyFile)
	if err != nil {
		return "", err
	}

	cfg, err := google.JWTConfigFromJSON(saKey)
	if err != nil {
		return "", err
	}

	expires := time.Now().Add(time.Second * 60)

	url, err = storage.SignedURL(bucket, filename, &storage.SignedURLOptions{
		GoogleAccessID: cfg.Email,
		PrivateKey:     cfg.PrivateKey,
		Method:         "GET",
		Expires:        expires,
	})

	return url, err
}
