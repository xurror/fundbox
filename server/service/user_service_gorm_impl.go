package service

// import (
// 	"context"
// 	"getting-to-go/model"

// 	"github.com/google/uuid"
// 	"github.com/sirupsen/logrus"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/clause"
// )

// // UserService provides user-related services
// type userServiceGormImpl struct {
// 	db     *gorm.DB
// 	logger *logrus.Logger
// }

// // NewUserService creates a new UserService instance
// func NewUserServiceGorm(logger *logrus.Logger, db *gorm.DB) UserService {
// 	return &userServiceGormImpl{
// 		db:     db,
// 		logger: logger,
// 	}
// }

// // CreateUser creates a new user
// func (s *userServiceGormImpl) CreateUser(ctx context.Context, firstName, lastName, email string) (*model.User, error) {
// 	user := &model.User{
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		Email:     email,
// 	}
// 	result := s.db.Create(&user)
// 	return user, model.HandleError(result.Error)
// }

// func (s *userServiceGormImpl) GetUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
// 	user := &model.User{}
// 	result := s.db.Preload(clause.Associations).First(&user, id)
// 	return user, model.HandleError(result.Error)
// }

// func (s *userServiceGormImpl) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
// 	user := &model.User{}
// 	result := s.db.Preload(clause.Associations).First(&user, "email = ?", email)
// 	return user, model.HandleError(result.Error)
// }

// func (s *userServiceGormImpl) GetUsers(
// 	ctx context.Context,
// 	limit int,
// 	offset interface{},
// ) ([]*model.User, interface{}, error) {
// 	var offsetVal int
// 	if offset == nil {
// 		offsetVal = 0
// 	} else {
// 		offsetVal = offset.(int)
// 	}

// 	var users []*model.User
// 	result := s.db.Preload(clause.Associations).
// 		Limit(limit).
// 		Offset(offsetVal).
// 		Find(&users)
// 	return users, nil, model.HandleError(result.Error)
// }

// // func (s *UserServiceGorm) GetUserContributions(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*model.Contribution, error) {
// // 	var contributions []*model.Contribution
// // 	result := s.db.
// // 		Preload(clause.Associations).
// // 		Limit(limit).
// // 		Offset(offset).
// // 		Find(&contributions, "contributor_id = ?", userID)
// // 	return contributions, model.HandleError(result.Error)
// // }
