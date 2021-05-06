package svc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SQSInput(msg string, url, key, val string) *sqs.SendMessageInput {
	return &sqs.SendMessageInput{
		MessageBody:  aws.String(msg),
		QueueUrl:     aws.String(url),
		DelaySeconds: aws.Int64(1),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			key: {
				DataType:    aws.String("String"),
				StringValue: aws.String(val),
			},
		},
	}
}
