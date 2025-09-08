package repository

import "github.com/rupayan-ninety-eight/tasker/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
