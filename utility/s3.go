package utility

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// User base structure for CP user
type User struct {
	AccountID  string
	UserID     string
	AccountUID string
	Groups     []int64
}

// Counts holds details about the uploaded file in S3
// like total number of valid and invalid items.
type Counts struct {
	TotalPages       int
	Total            int
	ValidEntries     int
	InvalidEntries   int
	DuplicateEntries int
}

// Contacts is a contact slice/array
type Contacts []*Contact

// Contact is a simple model that holds a basic contact details like mobile
// number, country, groups etc...
type Contact struct {
	PhoneNumber string
	Country     string
	Groups      []int64
	Attributes  map[string]string
}

// S3Object base structure for the uploaded s3 object should extend
// User.
type S3Object struct {
	Key    string
	Bucket string
	User
}

// Payload data struct
type Payload struct {
	S3Object
	SocketID string
	QueueURL string
}

// ImportPayload data struct for SNS payload
// from the API right after S3 upload
type ImportPayload struct {
	Payload
	FileID      string
	FileStateID int
	Counts      Counts
	Contacts    Contacts
}

// ImportPayloadWithColumn ...
type ImportPayloadWithColumn struct {
	ImportPayload
	Column map[string]int
}

// S3GetFromKeyWithConfig...
func S3GetFromKeyWithConfig(key string, bucket string, cfg *aws.Config, decompress bool) ([]byte, error) {
	svc := s3.New(session.New(), cfg)

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	item, err := svc.GetObject(params)
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
