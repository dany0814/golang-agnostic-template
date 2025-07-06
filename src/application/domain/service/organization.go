package domain

import (
	"context"
)

type IOrganizationService interface {
	Register(ctx context.Context, user any) any
	Update(ctx context.Context, id string, user any) any
}

type OganizationService struct {
}

func NewOrganizationService() IOrganizationService {
	return &OganizationService{}
}

func (org *OganizationService) Register(ctx context.Context, user any) any {
	return true
}

func (org *OganizationService) Update(ctx context.Context, id string, user any) any {
	return true
}
