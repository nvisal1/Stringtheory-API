package SQS

import (
	. "Stringtheory-API/user-curriculum-progress/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"reflect"
)

type SQSMessageStore struct {
	client *sqs.SQS
	queueURL string
}

func NewSQSMessageStore(queueURL string) (*SQSMessageStore, error) {
	session, err := session.NewSession(&aws.Config{
		Region:    aws.String("us-east-2"),
	})
    if err != nil {
        return nil, err
    }

    return &SQSMessageStore{client: sqs.New(session), queueURL: queueURL}, nil
}

func (sqsMessageStore SQSMessageStore) sendMessage(message *Message) error {
	formattedAttributes := make(map[string]*sqs.MessageAttributeValue)
	for key, value := range message.Attributes {
		formattedAttributes[key] = &sqs.MessageAttributeValue{
			DataType:         aws.String(reflect.TypeOf(value).String()),
			StringValue:      aws.String(value.(string)),
		}
	}

	_, err := sqsMessageStore.client.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: formattedAttributes,
		MessageBody: aws.String(message.Body),
		QueueUrl:    &sqsMessageStore.queueURL,
	})

	if err != nil {
		return err
	}

	return nil
}

func (sqsMessageStore SQSMessageStore) receiveMessage() (*Message, error) {

}


