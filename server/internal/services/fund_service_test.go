package services

import (
	"community-funds/internal/models"
	"community-funds/internal/repositories"
	"community-funds/internal/testutils"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestUser(t *testing.T, testDB *testutils.TestDatabase, number int) models.User {
	user := models.User{
		Auth0ID: strconv.Itoa(number),
		Name:    "Test User" + strconv.Itoa(number),
		Email:   "test" + strconv.Itoa(number) + "@example.com",
	}
	err := testDB.DB.Create(&user).Error
	require.NoError(t, err)
	return user
}

func TestCreateFund(t *testing.T) {
	testDB := testutils.SetupTestDB(t)
	fundRepo := repositories.NewFundRepository(testDB.DB)
	fundService := NewFundService(fundRepo)

	// Create test user
	user := createTestUser(t, testDB, 1)

	// Test fund creation
	fund, err := fundService.CreateFund("Test Fund", user.ID, 1000.0)
	require.NoError(t, err)
	assert.Equal(t, "Test Fund", fund.Name)
	assert.Equal(t, user.ID, fund.ManagerID)

	// Verify fund is in DB
	var dbFund models.Fund
	err = testDB.DB.First(&dbFund, "id = ?", fund.ID).Error
	require.NoError(t, err)
	assert.Equal(t, fund.ID, dbFund.ID)
}

func TestGetFund(t *testing.T) {
	testDB := testutils.SetupTestDB(t)

	fundRepo := repositories.NewFundRepository(testDB.DB)
	fundService := NewFundService(fundRepo)

	// Create test user
	user := createTestUser(t, testDB, 1)

	// Test fund creation
	testFund, err := fundService.CreateFund("Test Fund", user.ID, 1000.0)
	require.NoError(t, err)
	assert.Equal(t, "Test Fund", testFund.Name)
	assert.Equal(t, user.ID, testFund.ManagerID)

	// Get a fund
	fund, err := fundService.GetFund(testFund.ID)
	require.NoError(t, err)
	assert.Equal(t, "Test Fund", fund.Name)
	assert.Equal(t, user.ID, fund.ManagerID)
}

func TestGetFundsManagedByUser(t *testing.T) {
	testDB := testutils.SetupTestDB(t)

	fundRepo := repositories.NewFundRepository(testDB.DB)
	fundService := NewFundService(fundRepo)

	// Create test user
	user1 := createTestUser(t, testDB, 1)
	user2 := createTestUser(t, testDB, 2)

	// Create test funds
	testFunds := []struct {
		name         string
		managerID    uuid.UUID
		targetAmount float64
	}{
		{name: "Fund 1", managerID: user1.ID, targetAmount: 1000.0},
		{name: "Fund 3", managerID: user1.ID, targetAmount: 3000.0},
		{name: "Fund 2", managerID: user2.ID, targetAmount: 2000.0},
	}

	// Create test funds
	for _, fund := range testFunds {
		_, err := fundService.CreateFund(fund.name, fund.managerID, fund.targetAmount)
		require.NoError(t, err)
	}

	// Get funds managed by user21
	funds, err := fundService.GetFundsManagedByUser(&user1.ID)
	require.NoError(t, err)
	assert.Len(t, funds, 2)
	for i, fund := range funds {
		assert.Equal(t, testFunds[i].name, fund.Name)
		assert.Equal(t, testFunds[i].targetAmount, fund.TargetAmount)
		assert.Equal(t, user1.ID, fund.ManagerID)
	}

	// Get funds managed by user2
	funds, err = fundService.GetFundsManagedByUser(&user2.ID)
	require.NoError(t, err)
	assert.Len(t, funds, 1)
	assert.Equal(t, "Fund 2", funds[0].Name)
	assert.Equal(t, 2000.0, funds[0].TargetAmount)
	assert.Equal(t, user2.ID, funds[0].ManagerID)
}

func TestGetContributedFunds(t *testing.T) {
	testDB := testutils.SetupTestDB(t)

	fundRepo := repositories.NewFundRepository(testDB.DB)
	fundService := NewFundService(fundRepo)
	contributionService := NewContributionService(repositories.NewContributionRepository(testDB.DB))

	// Create test users
	fundManager := createTestUser(t, testDB, 1)
	contributor := createTestUser(t, testDB, 2)

	// Create test funds
	testFunds := []struct {
		name         string
		managerID    uuid.UUID
		targetAmount float64
	}{
		{name: "Fund 1", managerID: fundManager.ID, targetAmount: 1000.0},
		{name: "Fund 2", managerID: fundManager.ID, targetAmount: 2000.0},
	}

	// Create test funds
	for _, fund := range testFunds {
		createFund, err := fundService.CreateFund(fund.name, fund.managerID, fund.targetAmount)
		contributionService.MakeContribution(createFund.ID, &contributor.ID, 100.0, false)
		require.NoError(t, err)
	}

	// Get funds contributed to by user1
	contributedFunds, err := fundService.GetContributedFunds(fundManager.ID)
	require.NoError(t, err)
	assert.Len(t, contributedFunds, 0)

	// Get funds contributed to by user2
	contributedFunds, err = fundService.GetContributedFunds(contributor.ID)
	require.NoError(t, err)
	assert.Len(t, contributedFunds, 2)
}
