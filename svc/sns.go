package svc

import (
	"os"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSPayload struct {
	Message          string `json:"Message"`
	MessageId        string `json:"MessageId"`
	Signature        string `json:"Signature"`
	SignatureVersion string `json:"SignatureVersion"`
	SigningCertURL   string `json:"SigningCertURL"`
	SubscribeURL     string `json:"SubscribeURL"`
	Subject          string `json:"Subject"`
	Timestamp        string `json:"Timestamp"`
	Token            string `json:"Token"`
	TopicArn         string `json:"TopicArn"`
	Type             string `json:"Type"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
}

func GetSNSPayload(req []byte) (event *SNSPayload) {
	err := json.Unmarshal(req, &event)
	if err != nil {
		return nil
	}

	return event
}

func SNSMessage(message []byte) error {
	topicArn := os.Getenv("SNS_TOPIC")

	cfg := GetConfig()
	snsSVC := sns.New(session.New(), cfg)

	input := &sns.PublishInput{
		Message:  aws.String(string(message)),
		TopicArn: aws.String(topicArn),
	}

	_, err := snsSVC.Publish(input)
	return err
}
