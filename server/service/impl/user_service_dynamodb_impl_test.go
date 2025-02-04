package service

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"getting-to-go/model"
	"getting-to-go/service/mocks"
)

func marshalUser(user model.User) map[string]types.AttributeValue {
	userItem, err := attributevalue.MarshalMap(user)
	if err != nil {
		panic(err)
	}
	return userItem
}

// func TestCreateUsersTable(t *testing.T) {
// 	logger := logrus.New()
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockClient := mocks.NewMockDynamoDBClient(ctrl)
// 	svc := NewUserServiceDynamoDbImpl(logger, mockClient)
// 	ctx := context.TODO()

// 	mockClient.EXPECT().CreateTable(gomock.Any(), gomock.Any()).Return(&dynamodb.CreateTableOutput{}, nil)
// 	mockClient.EXPECT().DescribeTable(gomock.Any(), gomock.Any()).Return(&dynamodb.DescribeTableOutput{
// 		Table: &types.TableDescription{},
// 	}, nil)

// 	tableDesc, err := svc.createUsersTable(ctx)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, tableDesc)
// }

func TestCreateUser(t *testing.T) {
	logger := logrus.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockDynamoDBClient(ctrl)
	svc := NewUserServiceDynamoDbImpl(logger, mockClient)
	ctx := context.TODO()
	auth0Id := "auth0|12345"

	mockClient.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(&dynamodb.PutItemOutput{}, nil)

	user, err := svc.CreateUser(ctx, auth0Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestCreateUser_Error(t *testing.T) {
	logger := logrus.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockDynamoDBClient(ctrl)
	svc := NewUserServiceDynamoDbImpl(logger, mockClient)
	ctx := context.TODO()
	auth0Id := ""

	// mockClient.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

	user, err := svc.CreateUser(ctx, auth0Id)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestGetUserById_NotFound(t *testing.T) {
	logger := logrus.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockDynamoDBClient(ctrl)
	svc := NewUserServiceDynamoDbImpl(logger, mockClient)
	ctx := context.TODO()

	id := uuid.New()
	mockClient.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{
		Item: nil,
	}, nil)

	user, err := svc.GetUserById(ctx, id)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestGetUserByAuth0Id(t *testing.T) {
	logger := logrus.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockDynamoDBClient(ctrl)
	svc := NewUserServiceDynamoDbImpl(logger, mockClient)
	ctx := context.TODO()

	auth0Id := "auth0|12345"

	mockResponse := marshalUser(*model.NewUser(uuid.New(), auth0Id))

	mockClient.EXPECT().Query(gomock.Any(), gomock.Any()).Return(&dynamodb.QueryOutput{
		Items: []map[string]types.AttributeValue{mockResponse},
	}, nil)

	user, err := svc.GetUserByAuth0Id(ctx, auth0Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, auth0Id, user.Auth0Id)
}

func TestGetUsers_WithOffset(t *testing.T) {
	logger := logrus.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockDynamoDBClient(ctrl)
	svc := NewUserServiceDynamoDbImpl(logger, mockClient)
	ctx := context.TODO()

	mockUsers := []map[string]types.AttributeValue{
		marshalUser(*model.NewUser(uuid.New(), "auth0|12345")),
		marshalUser(*model.NewUser(uuid.New(), "auth0|67890")),
	}

	mockClient.EXPECT().Scan(gomock.Any(), gomock.Any()).Return(&dynamodb.ScanOutput{
		Items: mockUsers,
	}, nil)

	offset := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: uuid.New().String()},
	}

	users, lastEvaluatedKey, err := svc.GetUsers(ctx, 10, offset)

	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Nil(t, lastEvaluatedKey)
	assert.Len(t, users, 2)
}
func TestGetUserById(t *testing.T) {
	logger := logrus.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockDynamoDBClient(ctrl)
	svc := NewUserServiceDynamoDbImpl(logger, mockClient)
	ctx := context.TODO()

	id := uuid.New()
	mockResponse := marshalUser(*model.NewUser(id, "auth0|12345"))

	mockClient.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{
		Item: mockResponse,
	}, nil)

	user, err := svc.GetUserById(ctx, id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, id, user.Id)
}

func TestGetUserByAuth0Id_NotFound(t *testing.T) {
	logger := logrus.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockDynamoDBClient(ctrl)
	svc := NewUserServiceDynamoDbImpl(logger, mockClient)
	ctx := context.TODO()

	auth0Id := "auth0|nonexistent"
	mockClient.EXPECT().Query(gomock.Any(), gomock.Any()).Return(&dynamodb.QueryOutput{
		Items: []map[string]types.AttributeValue{},
	}, nil)

	user, err := svc.GetUserByAuth0Id(ctx, auth0Id)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestGetUsers(t *testing.T) {
	logger := logrus.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockDynamoDBClient(ctrl)
	svc := NewUserServiceDynamoDbImpl(logger, mockClient)
	ctx := context.TODO()

	mockUsers := []map[string]types.AttributeValue{
		marshalUser(*model.NewUser(uuid.New(), "auth0|12345")),
		marshalUser(*model.NewUser(uuid.New(), "auth0|67890")),
	}

	mockClient.EXPECT().Scan(gomock.Any(), gomock.Any()).Return(&dynamodb.ScanOutput{
		Items: mockUsers,
	}, nil)

	users, lastEvaluatedKey, err := svc.GetUsers(ctx, 10, nil)

	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Nil(t, lastEvaluatedKey)
	assert.Len(t, users, 2)
}
