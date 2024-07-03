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

type SQSProducer struct {
    client   *sqs.Client
    queueURL string
}

func NewSQSProducer(queueURL string) (*SQSProducer, error) {
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

	return &SQSProducer{
		client:   client,
		queueURL: queueURL,
	}, nil
}

func (p *SQSProducer) Publish(exchange, routingKey string, body []byte) error {
    _, err := p.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
        QueueUrl:    aws.String(p.queueURL),
        MessageBody: aws.String(string(body)),
    })
	fmt.Println("Message sent to queue")
    return err
}
