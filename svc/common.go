package svc

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func LogIt(message string) {
	logMessage := fmt.Sprintf("%s - %s", time.Now().Format(time.RFC3339), message)
	fmt.Fprintln(os.Stderr, logMessage)
}

func GetConfig() *aws.Config {
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ID"), os.Getenv("AWS_KEY"), "")
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return nil
	}

	return aws.NewConfig().WithRegion(os.Getenv("REGION")).WithCredentials(creds)

}

func GetAPISecret(secretName string) (secretBytes []byte, err error) {
	secretBytes, err = ioutil.ReadFile("/var/openfaas/secrets/" + secretName)

	return secretBytes, err
}

// DecompressBas64 decompresses base64 encoded and zlib compressed
// data.
func DecompressBas64(data []byte) (io.ReadCloser, error) {
	d, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	body, err := Decompress(d, "")
	return body, err
}

// Decompress decompresses data from the given compression.
// which defaults to zlib.
func Decompress(data []byte, compression string) (io.ReadCloser, error) {
	b := bytes.NewReader(data)
	var r io.ReadCloser
	var err error

	switch compression {
	case "gzip":
		r, err = gzip.NewReader(b)
	default:
		r, err = zlib.NewReader(b)
	}
	_ = r.Close()

	return r, err
}
