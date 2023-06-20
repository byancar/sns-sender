package app

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	settings "sns-sender/internal"
)

func SendMessage(message string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	cfg.Region = "us-east-1"

	client := sns.NewFromConfig(cfg)

	params := &sns.PublishInput{
		Message: aws.String(message),

		TopicArn: aws.String(settings.Env.AWS.TopicName),
	}

	resp, err := client.Publish(context.TODO(), params)
	if err != nil {
		return err
	}

	fmt.Println("Message sent successfully. Message ID:", *resp.MessageId)
	return nil
}
