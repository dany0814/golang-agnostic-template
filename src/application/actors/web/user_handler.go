package web

import (
	"golang-agnostic-template/src/application/domain/dto"
	domain "golang-agnostic-template/src/application/domain/service"
	"golang-agnostic-template/src/application/domain/utils"
	"golang-agnostic-template/src/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService domain.IUserService
	logger      logger.ILogger
}

func NewUserHandler(domainUserService domain.IUserService, logger logger.ILogger) *UserHandler {
	return &UserHandler{
		userService: domainUserService,
		logger:      logger,
	}
}

func (h *UserHandler) RegisterUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var registerUserReq dto.RegisterUserReq

		if err := ctx.BindJSON(&registerUserReq); err != nil {
			h.logger.Error(utils.HANDLER+utils.USER, logger.LoggerField{Key: "BindJSON", Value: err})
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "REQUEST", Value: registerUserReq})

		res, err := h.userService.Register(ctx, registerUserReq)
		if err != nil {
			h.logger.Error(utils.HANDLER+utils.USER, logger.LoggerField{Key: "Register", Value: err})
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "REQUEST", Value: res})
		ctx.JSON(http.StatusCreated, res)
	}
}

func (h *UserHandler) LoginUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginUserReq dto.LoginUserReq

		if err := ctx.BindJSON(&loginUserReq); err != nil {
			h.logger.Error(utils.HANDLER+utils.USER, logger.LoggerField{Key: "BindJSON", Value: err})
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "REQUEST", Value: loginUserReq})

		res, err := h.userService.Login(ctx, loginUserReq)
		if err != nil {
			h.logger.Error(utils.HANDLER+utils.USER, logger.LoggerField{Key: "Register", Value: err})
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "REQUEST", Value: res})
		ctx.JSON(http.StatusAccepted, res)
	}
}

func (h *UserHandler) GetUserByIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "PARAM", Value: id})

		res, err := h.userService.GetUserById(ctx, id)
		if err != nil {
			h.logger.Error(utils.HANDLER+utils.USER, logger.LoggerField{Key: "Register", Value: err})
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "REQUEST", Value: res})
		ctx.JSON(http.StatusAccepted, res)
	}
}

func (h *UserHandler) UpdateUserByIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updateUserReq dto.UpdateUserReq

		if err := ctx.BindJSON(&updateUserReq); err != nil {
			h.logger.Error(utils.HANDLER+utils.USER, logger.LoggerField{Key: "BindJSON", Value: err})
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "REQUEST", Value: updateUserReq})

		id := ctx.Param("id")
		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "PARAM", Value: id})

		res, err := h.userService.UpdateUserById(ctx, id, updateUserReq)
		if err != nil {
			h.logger.Error(utils.HANDLER+utils.USER, logger.LoggerField{Key: "Register", Value: err})
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "REQUEST", Value: res})
		ctx.JSON(http.StatusAccepted, res)
	}
}

func (h *UserHandler) DeleteUserByIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "PARAM", Value: id})

		res, err := h.userService.DeleteUserById(ctx, id)
		if err != nil {
			h.logger.Error(utils.HANDLER+utils.USER, logger.LoggerField{Key: "Register", Value: err})
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "REQUEST", Value: res})
		ctx.JSON(http.StatusAccepted, res)
	}
}
