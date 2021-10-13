package svc

import (
	"encoding/json"
	"os"

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

type AttributesKey struct {
	Name        string
	DataType    string
	StringValue string
	AWSCred     bool `json:"AWSCred" default:"false"`
}

func GetSNSPayload(req []byte) (event *SNSPayload) {
	err := json.Unmarshal(req, &event)
	if err != nil {
		return nil
	}

	return event
}

func SNSMessage(message []byte, msgAttrKey *AttributesKey) error {
	topicArn := os.Getenv("SNS_TOPIC")

	var sess *session.Session
	if msgAttrKey.AWSCred {
		sess, _ = session.NewSession(&aws.Config{})
	} else {
		cfg := GetConfig()
		sess, _ = session.NewSession(cfg)
	}

	snsSVC := sns.New(sess)

	input := &sns.PublishInput{
		Message:  aws.String(string(message)),
		TopicArn: aws.String(topicArn),
	}

	if msgAttrKey.DataType != "" && msgAttrKey.StringValue != "" {
		key := msgAttrKey.Name
		input = &sns.PublishInput{
			Message:  aws.String(string(message)),
			TopicArn: aws.String(topicArn),
			MessageAttributes: map[string]*sns.MessageAttributeValue{
				key: {
					DataType:    aws.String(msgAttrKey.DataType),
					StringValue: aws.String(msgAttrKey.StringValue),
				},
			},
		}
	}

	_, err := snsSVC.Publish(input)
	return err
}
