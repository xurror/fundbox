package model

import (
	"context"
	appConfig "getting-to-go/config"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDBClient(c *appConfig.AppConfig) *dynamodb.Client {
	provider := aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     c.Aws.Credentials.AccessKey,
			SecretAccessKey: c.Aws.Credentials.SecretKey,
		}, nil
	})

	// Configure AWS SDK
	cfg, err := awsConfig.LoadDefaultConfig(
		context.Background(),
		awsConfig.WithRegion(c.Aws.Region),
		awsConfig.WithCredentialsProvider(provider),
	)
	if err != nil {
		log.Fatal(err)
	}

	return dynamodb.NewFromConfig(cfg)
}
