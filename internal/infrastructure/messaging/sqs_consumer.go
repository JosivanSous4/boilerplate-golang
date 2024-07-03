package messaging

import (
	configLocal "boilerplate-go/internal/config"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSConsumer struct {
	client   *sqs.Client
	queueURL string
}

func NewSQSConsumer(queueURL string) (*SQSConsumer, error) {
	cfgLocal := configLocal.LoadConfig()

	awsEndpoint := cfgLocal.AwsEndpoint
	awsRegion := cfgLocal.AwsRegion

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		})),
	)
	if err != nil {
		log.Fatalf("failed to load AWS SDK config: %v", err)
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	return &SQSConsumer{
		client:   client,
		queueURL: queueURL,
	}, nil
}

func (c *SQSConsumer) Consume(queueName string) (<-chan []byte, error) {
	out := make(chan []byte)
	go func() {
		for {
			result, err := c.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
				QueueUrl:            aws.String(c.queueURL),
				MaxNumberOfMessages: 10,
				WaitTimeSeconds:     10,
			})

            if err != nil {
				log.Printf("failed to fetch sqs message %v", err)
				continue
			}

			for _, message := range result.Messages {
				fmt.Println("Message read from queue")

				out <- []byte(*message.Body)
				_, err := c.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(c.queueURL),
					ReceiptHandle: message.ReceiptHandle,
				})
				if err != nil {
					log.Printf("failed to delete sqs message %v", err)
				}
			}
		}
	}()
	return out, nil
}
