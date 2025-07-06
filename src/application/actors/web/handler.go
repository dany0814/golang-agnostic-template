package web

import (
	"golang-agnostic-template/src/application/domain/dto"
	domain "golang-agnostic-template/src/application/domain/service"
	"golang-agnostic-template/src/application/domain/utils"
	"golang-agnostic-template/src/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService domain.IUserService
	orgService  domain.IOrganizationService
	logger      logger.ILogger
}

func NewHandler(domainUserService domain.IUserService, domainOrgService domain.IOrganizationService, logger logger.ILogger) *Handler {
	return &Handler{
		userService: domainUserService,
		orgService:  domainOrgService,
		logger:      logger,
	}
}

func (h *Handler) RegisterUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var registerUserReq dto.RegisterUserReq

		if err := ctx.BindJSON(&registerUserReq); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "request", Value: registerUserReq})

		res, err := h.userService.Register(ctx, registerUserReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		h.logger.Info(utils.HANDLER+utils.USER, logger.LoggerField{Key: "response", Value: res})
		ctx.JSON(http.StatusCreated, res)
	}
}

func (h *Handler) RegisterOrganizationHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, "oka")
	}
}
