package service

import (
	"context"
	"time"

	"github.com/shereifsrf/savespent-api/dao"
	"github.com/shereifsrf/savespent-api/dao/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// add session service interface
type SessionService interface {
	AddSession(bSession *BindedSession) (*model.UserSession, error)
	addUserSession(bSession *BindedSession) (*model.UserSession, error)
}

// add session service struct
type sessoinService struct {
	ctx         context.Context
	userService UserService
}

type BindedSession struct {
	UserID     string `json:"user_id" binding:"required"`
	Email      string `json:"email" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	DeviceType string `json:"device_type" binding:"required"`
}

func NewSessionService(ctx context.Context) SessionService {
	return &sessoinService{
		ctx:         ctx,
		userService: NewUserService(ctx),
	}
}

func (s *sessoinService) AddSession(bSession *BindedSession) (*model.UserSession, error) {
	// check if user exist
	user := s.userService.userExist(bSession.UserID)
	if user == nil {
		if _, err := s.userService.AddUser(bSession); err != nil {
			return nil, err
		}
	}

	// add user session
	session, err := s.addUserSession(bSession)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// add a record in mongodb user session
func (s *sessoinService) addUserSession(session *BindedSession) (*model.UserSession, error) {
	// add user session
	userSession := &model.UserSession{
		ID:          primitive.NewObjectID(),
		UserID:      session.UserID,
		DeviceType:  session.DeviceType,
		CreatedDate: time.Now(),
	}

	if _, err := dao.UsMongo.InsertOne(s.ctx, userSession); err != nil {
		return nil, err
	}

	return userSession, nil
}
