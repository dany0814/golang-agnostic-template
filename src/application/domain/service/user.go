package domain

import (
	"context"
	"golang-agnostic-template/src/application/domain/business"
	"golang-agnostic-template/src/application/domain/dto"
	entity "golang-agnostic-template/src/application/domain/model"
	"golang-agnostic-template/src/application/domain/repository"
	"golang-agnostic-template/src/application/domain/utils"
)

type IUserService interface {
	Register(ctx context.Context, user dto.RegisterUserReq) (res dto.RegisterUserRes, err error)
}

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (usr *UserService) Register(ctx context.Context, user dto.RegisterUserReq) (res dto.RegisterUserRes, err error) {
	var registerUserRes dto.RegisterUserRes
	pass, err := business.HashAndSalt(user.Password)
	if err != nil {
		return registerUserRes, err
	}
	cellphone, err := utils.IsValidPhone(user.Phone)
	userModel := entity.User{
		Email:        user.Email,
		Username:     user.UserName,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Phone:        cellphone,
		Password:     pass,
		Organization: user.OrganizationDomain,
	}
	err = userModel.ValidateEmail()
	if err != nil {
		return registerUserRes, err
	}
	userModel.BuildUser()

	userCreated, err := usr.userRepository.Create(ctx, userModel)
	obfuscatePhone, err := utils.ObfuscatePhoneNumber(user.Phone)

	registerUserRes = dto.RegisterUserRes{
		Email:        userCreated.Email,
		Password:     utils.HASH_PASS,
		Phone:        obfuscatePhone,
		FirstName:    userCreated.FirstName,
		LastName:     userCreated.LastName,
		UserName:     userCreated.Username,
		Organization: userCreated.Organization,
	}
	return registerUserRes, nil
}
