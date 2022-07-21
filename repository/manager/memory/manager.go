package memory

import (
	"github.com/SebastienLaurent/structure/repository"
	"github.com/SebastienLaurent/structure/repository/git"
)

type Manager struct {	
}

func New() (*Manager) {
	return &Manager{}
}

func (m *Manager) Get(url string) (repository.Repository, error)  {
	// Simple implementation that always create a in memory repo
	return git.NewMemoryRepository(url)
}