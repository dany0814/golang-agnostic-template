package domain

import (
	"context"
	"errors"
	"golang-agnostic-template/src/application/domain/business"
	"golang-agnostic-template/src/application/domain/dto"
	entity "golang-agnostic-template/src/application/domain/model"
	"golang-agnostic-template/src/application/domain/repository"
	"golang-agnostic-template/src/application/domain/utils"
	"golang-agnostic-template/src/pkg/jwt"
	"golang-agnostic-template/src/pkg/logger"
)

type IUserService interface {
	Register(ctx context.Context, user dto.RegisterUserReq) (res dto.RegisterUserRes, err error)
	Login(ctx context.Context, user dto.LoginUserReq) (res dto.LoginUserRes, err error)
	GetUserById(ctx context.Context, userId string) (res dto.GetUserRes, err error)
	UpdateUserById(ctx context.Context, userId string, user dto.UpdateUserReq) (res dto.UpdateUserRes, err error)
	DeleteUserById(ctx context.Context, userId string) (res dto.DeleteUserRes, err error)
}

type UserService struct {
	userRepository repository.IUserRepository
	logger         logger.ILogger
}

func NewUserService(userRepository repository.IUserRepository, logger logger.ILogger) IUserService {
	return &UserService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (usr *UserService) Register(ctx context.Context, user dto.RegisterUserReq) (res dto.RegisterUserRes, err error) {
	var registerUserRes dto.RegisterUserRes
	pass, err := business.HashAndSalt(user.Password)
	if err != nil {
		return registerUserRes, err
	}
	cellphone, err := business.IsValidPhone(user.Phone)
	userModel := entity.User{
		Username:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  pass,
		Phone:     cellphone,
		Settings: entity.UserSettings{
			EmailNotifications: user.EmailNotifications,
			SmsNotifications:   user.SmsNotifications,
			Language:           user.Language,
		},
	}
	err = userModel.ValidateEmail()
	if err != nil {
		return registerUserRes, err
	}
	userModel.BuildUser()
	userCreated, err := usr.userRepository.Create(ctx, &userModel)
	if err != nil {
		usr.logger.Error(utils.SERVICE+utils.USER, logger.LoggerField{Key: "RESPONSE", Value: err})
		return registerUserRes, err
	}
	obfuscatePhone, err := business.ObfuscatePhoneNumber(user.Phone)

	registerUserRes = dto.RegisterUserRes{
		ID:       userCreated.ID.String(),
		Email:    userCreated.Email,
		Password: utils.HASH_PASS,
		Phone:    obfuscatePhone,
		UserName: userCreated.Username,
		Language: userCreated.Settings.Language,
	}
	return registerUserRes, err
}

func (usr *UserService) Login(ctx context.Context, user dto.LoginUserReq) (res dto.LoginUserRes, err error) {
	var registerUserRes dto.LoginUserRes
	userFounded, err := usr.userRepository.Read(ctx, &user.Email, nil, nil)

	if err != nil {
		return registerUserRes, errors.New(err.Error())
	}
	if !business.ComparePasswords(userFounded.Password, user.Password) {
		return registerUserRes, errors.New(utils.ErrMsgIncorrectPassword)
	}

	idString, ok := userFounded.ID.ID.(string)
	if !ok {
		return registerUserRes, errors.New("No se pudo convertir ID a string")
	}

	token, err := jwt.CreateToken(&idString)
	if err != nil {
		return registerUserRes, errors.New(err.Error())
	}
	obfuscatePhone, err := business.ObfuscatePhoneNumber(userFounded.Phone)

	registerUserRes = dto.LoginUserRes{
		ID:       userFounded.ID.ID.(string),
		Email:    userFounded.Email,
		Password: utils.HASH_PASS,
		Phone:    obfuscatePhone,
		UserName: userFounded.Username,
		Language: userFounded.Settings.Language,
		Token:    token,
	}
	return registerUserRes, err
}

func (usr *UserService) GetUserById(ctx context.Context, userId string) (res dto.GetUserRes, err error) {
	var userRes dto.GetUserRes
	userFounded, err := usr.userRepository.Read(ctx, nil, nil, &userId)
	if err != nil {
		return userRes, errors.New(err.Error())
	}
	obfuscatePhone, err := business.ObfuscatePhoneNumber(userFounded.Phone)
	userRes = dto.GetUserRes{
		ID:        userFounded.ID.ID.(string),
		Username:  userFounded.Username,
		FirstName: userFounded.FirstName,
		LastName:  userFounded.LastName,
		Email:     userFounded.Email,
		State:     userFounded.State,
		Password:  utils.HASH_PASS,
		Phone:     obfuscatePhone,
		CreatedAt: userFounded.CreatedAt,
		UpdatedAt: userFounded.UpdatedAt,
		DeletedAt: userFounded.DeletedAt,
	}
	return userRes, err
}

func (usr *UserService) UpdateUserById(ctx context.Context, userId string, user dto.UpdateUserReq) (res dto.UpdateUserRes, err error) {
	var userRes dto.UpdateUserRes

	userModel := entity.User{
		Username:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
		Settings: entity.UserSettings{
			EmailNotifications: user.EmailNotifications,
			SmsNotifications:   user.SmsNotifications,
			Language:           user.Language,
		},
	}
	userDeleted, err := usr.userRepository.Update(ctx, &userId, &userModel)
	if err != nil {
		return userRes, errors.New(err.Error())
	}
	userRes = dto.UpdateUserRes{
		ID: userDeleted.ID.ID.(string),
	}
	return userRes, err
}

func (usr *UserService) DeleteUserById(ctx context.Context, userId string) (res dto.DeleteUserRes, err error) {
	var userRes dto.DeleteUserRes
	userDeleted, err := usr.userRepository.Delete(ctx, &userId)
	if err != nil {
		return userRes, errors.New(err.Error())
	}
	userRes = dto.DeleteUserRes{
		ID: userDeleted.ID.ID.(string),
	}
	return userRes, err
}
