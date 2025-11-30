package handler

import (
	"api-gateway/internal/service"
)

type Handler struct {
	service *service.ServiceRepositoryClient
}

func NewApiHandler(service *service.ServiceRepositoryClient) *Handler {
	return &Handler{
		service: service,
	}
}
