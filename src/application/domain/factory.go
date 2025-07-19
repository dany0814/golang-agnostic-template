package factory

import (
	"context"
	"fmt"
	surreal "golang-agnostic-template/src/application/actors/db"
	"golang-agnostic-template/src/application/domain/repository"
	domain "golang-agnostic-template/src/application/domain/service"
	"golang-agnostic-template/src/application/domain/utils"
	"golang-agnostic-template/src/pkg/database"
	"golang-agnostic-template/src/pkg/logger"
)

type FactoryRepository interface {
	Create(ctx context.Context) FactoryRepository
	GetUserRepository() repository.IUserRepository
}

type FactoryService interface {
	Create(ctx context.Context) FactoryService
	GetUserService() domain.IUserService
	GetOrganizationService() domain.IOrganizationService
}

type RepositoryFactory struct {
	db             database.IDatabaseConnection
	userRepository repository.IUserRepository
	logger         logger.ILogger
}

type ServiceFactory struct {
	UserService         domain.IUserService
	OrganizationService domain.IOrganizationService
	logger              logger.ILogger
}

func NewRepositoryFactory(db database.IDatabaseConnection, logger logger.ILogger) FactoryRepository {
	return &RepositoryFactory{
		db:     db,
		logger: logger,
	}
}

func NewServiceFactory(logger logger.ILogger) FactoryService {
	return &ServiceFactory{
		logger: logger,
	}
}

func (p *RepositoryFactory) Create(ctx context.Context) FactoryRepository {
	if p.db == nil {
		p.logger.Error(utils.ErrMsgDatabaseConnect, logger.LoggerField{Key: "error", Value: "database connection is nil"})
	}

	_, err := p.db.Connect(ctx)
	if err != nil {
		p.logger.Error(utils.ErrMsgDatabaseConnect,
			logger.LoggerField{Key: "error", Value: err})
	}

	p.userRepository = surreal.NewUserRepository(p.db)
	return p
}

func (p *RepositoryFactory) GetUserRepository() repository.IUserRepository {
	return p.userRepository
}

func (s *ServiceFactory) Create(ctx context.Context) FactoryService {
	repoFactory := NewRepositoryFactory(database.NewSurrealDBConnection(), s.logger)
	repoFactory = repoFactory.Create(ctx)
	s.UserService = domain.NewUserService(repoFactory.GetUserRepository(), s.logger)
	s.OrganizationService = domain.NewOrganizationService()
	return s
}

func (s *ServiceFactory) GetUserService() domain.IUserService {
	return s.UserService
}

func (s *ServiceFactory) GetOrganizationService() domain.IOrganizationService {
	return s.OrganizationService
}

type CustomPanicError struct {
	Message string
	Err     error
	Fields  []logger.LoggerField
}

func (e *CustomPanicError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}
