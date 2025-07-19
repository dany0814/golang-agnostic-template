package web

import (
	"context"
	factory "golang-agnostic-template/src/application/domain"
	"golang-agnostic-template/src/application/domain/utils"
	"golang-agnostic-template/src/pkg/logger"

	"github.com/gin-gonic/gin"
)

type GroupRoute struct {
	Name  string
	Paths []Route
}

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func NewGroupRoutes(ctx context.Context, log logger.ILogger) []GroupRoute {
	var routes []GroupRoute = make([]GroupRoute, 0)
	serviceFactory := factory.NewServiceFactory(log)
	services := Routes(serviceFactory, ctx, log)
	routes = append(routes, services...)
	return routes
}

func Routes(service factory.FactoryService, ctx context.Context, log logger.ILogger) []GroupRoute {
	srv := service.Create(ctx)
	userHandler := NewUserHandler(srv.GetUserService(), log)
	organizationHandler := NewOrganizationHandler(srv.GetOrganizationService(), log)
	routes := []GroupRoute{
		{Name: utils.USER,
			Paths: []Route{
				{Method: "POST", Path: "/create", Handler: userHandler.RegisterUserHandler()},
				{Method: "POST", Path: "/login", Handler: userHandler.LoginUserHandler()},
				{Method: "GET", Path: "/:id", Handler: userHandler.GetUserByIdHandler()},
				{Method: "PATCH", Path: "/:id", Handler: userHandler.UpdateUserByIdHandler()},
				{Method: "DELETE", Path: "/:id", Handler: userHandler.DeleteUserByIdHandler()},
			},
		},
		{Name: utils.ORGANIZATION,
			Paths: []Route{
				{Method: "POST", Path: "/create", Handler: organizationHandler.RegisterOrganizationHandler()},
				{Method: "GET", Path: "/:id", Handler: organizationHandler.RegisterOrganizationHandler()},
			},
		},
	}
	return routes
}
