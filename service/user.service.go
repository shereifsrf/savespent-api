package service

import (
	"context"
	"time"

	"github.com/shereifsrf/savespent-api/dao"
	"github.com/shereifsrf/savespent-api/dao/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	AddUser(bSession *BindedSession) (*model.User, error)
	userExist(id string) *model.User
}

type userService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) UserService {
	return &userService{
		ctx: ctx,
	}
}

// AddUser adds new user
// if user exist, then increment counter
func (u *userService) AddUser(bSession *BindedSession) (*model.User, error) {
	// best encryption for password
	// store the password in hashed format string
	passhash, err := bcrypt.GenerateFromPassword([]byte(bSession.UserID), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:          bSession.UserID,
		Email:       bSession.Email,
		FirstName:   bSession.FirstName,
		LastName:    bSession.LastName,
		Password:    string(passhash),
		CreatedDate: time.Now(),
	}
	if err := dao.MySql.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// check if user exist
func (u *userService) userExist(id string) *model.User {
	var user model.User
	if err := dao.MySql.Where("id = ?", id).First(&user).Error; err != nil {
		return nil
	}

	return &user
}
