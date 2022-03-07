package core

import "context"

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service
// Service interface
type Service interface {
	CanvasCreate(context.Context, *CanvasCreateRequest) (*CanvasCreateResponse, error)
	CanvasGet(context.Context, *CanvasGetRequest) (*CanvasGetResponse, error)
}

func New(c *Config, store Store) Service {
	return &canvasService{
		config: c,
		store:  store,
	}
}

type canvasService struct {
	config *Config
	store  Store
}
