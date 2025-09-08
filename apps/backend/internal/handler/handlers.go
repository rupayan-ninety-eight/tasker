package handler

import (
	"github.com/rupayan-ninety-eight/tasker/internal/server"
	"github.com/rupayan-ninety-eight/tasker/internal/service"
)

type Handlers struct {
	Health  *HealthHandler
	OpenAPI *OpenAPIHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:  NewHealthHandler(s),
		OpenAPI: NewOpenAPIHandler(s),
	}
}
