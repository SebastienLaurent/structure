package disk

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
	// Le manager gere un pool de clone de repo
	// Un meme repo (meme url de clone) peut etre present plusieurs fois
	// Un repo ne peux etre donné qu'une seule fois par un Get
	
	path := "computed"
	return git.NewDiskRepository(path,url)
}