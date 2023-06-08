package services

import (
	"context"
	"errors"
	"getting-to-go/graph/generated"
	"getting-to-go/models"
	"getting-to-go/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"time"
)

// UserService provides user-related services
type UserService struct{}

// NewUserService creates a new UserService instance
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(firstName, lastName, email, password string) (*models.User, error) {
	return models.CreateUser(&models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Roles:     models.ConvertToPQStringArray([]models.Role{"INITIATOR"}),
	})
}

func (s *UserService) CreateUserFromInput(input generated.NewUser) (*models.User, error) {
	return models.CreateUser(&models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  uuid.NewString(),
		Roles:     models.ConvertToPQStringArray(input.Roles),
	})
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id uuid.UUID) (*models.User, error) {
	return models.GetUser(id)
}

// GetUsers retrieves a list of users
func (s *UserService) GetUsers(limit, offset int) ([]*models.User, error) {
	return models.GetUsers(limit, offset)
}

// Authenticate authenticates a user by email and password
func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	// Get the user by email
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// If the user does not exist, return an error
	if user == nil {
		return nil, utils.NewError(http.StatusUnauthorized, "Invalid email or password")
	}

	// If the password does not match, return an error
	if !utils.CheckPasswordHash(user.Password, password) {
		return nil, utils.NewError(http.StatusUnauthorized, "Invalid email or password")
	}

	return user, nil
}

var validate = validator.New()

func (s *UserService) SignUp(newUser generated.NewUser) (*models.User, error) {
	// FIXME: This is a temporary hack to get the user service working
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	errList := gqlerror.List{}
	//var user models.User

	// FIXME: This is a temporary hack to get the user service working
	//if err := c.BindJSON(&user); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//validationErr := validate.Struct(user)
	//if validationErr != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
	//	return
	//}

	user, err := models.GetUserByEmail(newUser.Email)
	defer cancel()
	if err == nil {
		log.Panic(err)
		errList = append(errList, gqlerror.Wrap(errors.New("Email Already Exists")))

		//c.JSON(http.StatusInternalServerError, gin.H{"error": "error detected while fetching the email"})
	}

	password := utils.HashPassword(user.Password)
	user.Password = password

	//count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	defer cancel()
	// FIXME: This is a temporary hack to get the user service working
	//if err != nil {
	//	log.Panic(err)
	//	c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error occured while fetching the phone number"})
	//}
	//
	//if count > 0 {
	//	c.JSON(http.StatusInternalServerError, gin.H{"Error": "The mentioned E-Mail or Phone Number already exists"})
	//}

	user, err = s.CreateUser(
		newUser.FirstName,
		newUser.LastName,
		newUser.Email,
		password,
	)

	if err != nil {
		log.Panic(err)
		errList = append(errList, gqlerror.Wrap(errors.New("Error while creating user")))
		//c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error while creating user"})
	}

	// FIXME: This is a temporary hack to get the user service working
	//token, refreshToken, _ := utils.GenerateAllTokens(user.ID, user.Email)

	//resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	//if insertErr != nil {
	//	msg := fmt.Sprintf("User Details were not Saved")
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
	//	return
	//}
	defer cancel()
	//return token, refreshToken
	return user, errList
	//c.JSON(http.StatusOK, resultInsertionNumber)
}
