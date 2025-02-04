package service

import (
	"context"
	"errors"
	"getting-to-go/model"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// UserService provides user-related services
type userServiceDynamoDbImpl struct {
	logger         *logrus.Logger
	tableName      string
	dynamoDbClient model.DynamoDBClient
}

// NewUserService creates a new UserService instance
func NewUserServiceDynamoDbImpl(
	logger *logrus.Logger,
	dynamoDbClient model.DynamoDBClient,
) model.UserService {
	return &userServiceDynamoDbImpl{
		logger:         logger,
		tableName:      "Fundbox-Users",
		dynamoDbClient: dynamoDbClient,
	}
}

// CreateMovieTable creates a DynamoDB table with a composite primary key defined as
// a string sort key named `title`, and a numeric partition key named `year`.
// This function uses NewTableExistsWaiter to wait for the table to be created by
// DynamoDB before it returns.
func (s *userServiceDynamoDbImpl) createUsersTable(ctx context.Context) (*types.TableDescription, error) {
	var tableDesc *types.TableDescription
	table, err := s.dynamoDbClient.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(s.tableName),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeB,
			},
			{
				AttributeName: aws.String("email"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("email-index"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("email"),
						KeyType:       types.KeyTypeHash,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})
	if err != nil {
		s.logger.Debugf("Couldn't create table %v. Here's why: %v\n", s.tableName, err)
		return nil, err
	}

	waiter := dynamodb.NewTableExistsWaiter(s.dynamoDbClient)
	err = waiter.Wait(ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(s.tableName)}, 5*time.Minute)
	if err != nil {
		s.logger.Debugf("Wait for table exists failed. Here's why: %v\n", err)
	}
	tableDesc = table.TableDescription
	return tableDesc, err
}

// CreateUser creates a new user
func (s *userServiceDynamoDbImpl) CreateUser(
	ctx context.Context,
	auth0Id string,
) (*model.User, error) {
	if auth0Id == "" {
		return nil, errors.New("auth0Id cannot be empty")
	}

	user := model.NewUser(uuid.New(), auth0Id)
	userItem, err := attributevalue.MarshalMap(user)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	s.logger.Infof("Item to Commit: %v\n", user)

	_, err = s.dynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(s.tableName),
		Item:      userItem,
	})
	if err != nil {
		s.logger.Debugf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}

	return user, err
}

// GetUser retrieves a user by ID
func (s *userServiceDynamoDbImpl) GetUserById(
	ctx context.Context,
	id uuid.UUID,
) (*model.User, error) {
	resp, err := s.dynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(s.tableName),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id.String()},
		},
	})
	if err != nil {
		return nil, err
	}

	if resp.Item == nil {
		return nil, model.ErrNotFound
	}

	user := &model.User{}
	err = attributevalue.UnmarshalMap(resp.Item, user)
	if err != nil {
		s.logger.Debugf("Couldn't unmarshal response. Here's why: %v\n", err)
		return nil, err
	}
	return user, nil
}

// GetUser retrieves a user by ID
func (s *userServiceDynamoDbImpl) GetUserByAuth0Id(
	ctx context.Context,
	auth0Id string,
) (*model.User, error) {
	resp, err := s.dynamoDbClient.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(s.tableName),
		IndexName:              aws.String("auth0_id-index"),
		KeyConditionExpression: aws.String("auth0_id = :auth0_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":auth0_id": &types.AttributeValueMemberS{Value: auth0Id},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(resp.Items) == 0 {
		return nil, model.ErrNotFound
	}

	user := &model.User{}
	err = attributevalue.UnmarshalMap(resp.Items[0], user)
	if err != nil {
		s.logger.Debugf("Couldn't unmarshal response. Here's why: %v\n", err)
		return nil, err
	}
	return user, nil
}

// GetUsers retrieves a paginated list of users
func (s *userServiceDynamoDbImpl) GetUsers(
	ctx context.Context,
	limit int,
	offset interface{},
) ([]*model.User, interface{}, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(s.tableName),
		Limit:     aws.Int32(int32(limit)),
	}
	if offset != nil {
		scanInput.ExclusiveStartKey = offset.(map[string]types.AttributeValue)
	}

	resp, err := s.dynamoDbClient.Scan(ctx, scanInput)
	if err != nil {
		return nil, nil, err
	}

	var users []*model.User
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &users)
	if err != nil {
		s.logger.Debugf("Couldn't unmarshal query response. Here's why: %v\n", err)
		return nil, nil, err
	}
	return users, resp.LastEvaluatedKey, nil
}
