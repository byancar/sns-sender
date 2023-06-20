package app

import (
	"context"
	"errors"
	settings "sns-sender/internal"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	_ "github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/stretchr/testify/assert"
)

type MockSNSClient struct {
	sns.Client
	mockPublish func(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

func (m *MockSNSClient) Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {
	return m.mockPublish(ctx, params, optFns...)
}

type MockConfigLoader struct {
	mockLoadDefaultConfig func(ctx context.Context) (aws.Config, error)
}

func (m *MockConfigLoader) LoadDefaultConfig(ctx context.Context) (aws.Config, error) {
	return m.mockLoadDefaultConfig(ctx)
}

func TestSendMessage_Success(t *testing.T) {
	message := "Hello, world!"
	mockClient := &MockSNSClient{}

	expectedParams := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(settings.Env.AWS.TopicName),
	}
	mockClient.mockPublish = func(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {
		assert.Equal(t, expectedParams, params)
		return &sns.PublishOutput{
			MessageId: aws.String("123"),
		}, nil
	}

	err := SendMessage(message)

	assert.NoError(t, err)
}

func TestSendMessage_LoadConfigError(t *testing.T) {
	message := "Hello, world!"
	mockLoader := &MockConfigLoader{}

	mockLoader.mockLoadDefaultConfig = func(ctx context.Context) (aws.Config, error) {
		return aws.Config{}, errors.New("failed to load config")
	}

	err := SendMessage(message)

	assert.Error(t, err)
}
