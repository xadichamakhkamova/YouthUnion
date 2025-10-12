package handler

import (
	"api-gateway/internal/service"
	"context"
)

type Handler struct {
	service *service.ServiceRepositoryClient
}

func NewApiHandler(service *service.ServiceRepositoryClient) *Handler {
	return &Handler{
		service: service,
	}
}

var ctx = context.Background()
