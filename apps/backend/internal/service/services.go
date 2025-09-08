package service

import (
	"github.com/rupayan-ninety-eight/tasker/internal/lib/job"
	"github.com/rupayan-ninety-eight/tasker/internal/repository"
	"github.com/rupayan-ninety-eight/tasker/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
