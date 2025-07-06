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
	handler := NewHandler(srv.GetUserService(), srv.GetOrganizationService(), log)
	routes := []GroupRoute{
		{Name: utils.USER,
			Paths: []Route{
				{Method: "POST", Path: "/create", Handler: handler.RegisterUserHandler()},
				{Method: "GET", Path: "/:id", Handler: handler.RegisterUserHandler()},
			},
		},
		{Name: utils.ORGANIZATION,
			Paths: []Route{
				{Method: "POST", Path: "/create", Handler: handler.RegisterUserHandler()},
				{Method: "GET", Path: "/:id", Handler: handler.RegisterUserHandler()},
			},
		},
	}
	return routes
}
