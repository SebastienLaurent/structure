package git

import(
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-billy/v5"
)

type Repository struct {
	repository *gogit.Repository
}

func NewDiskRepository(path string, url string) (*Repository, error) {
	return nil, nil
}

func NewMemoryRepository(url string) (*Repository, error) {
	return nil, nil
}

func (r *Repository) Add() {
}

func (r *Repository) Commit() {
}

func (r *Repository) Clean() {
}

func (r *Repository) Pull() {
}

func (r *Repository) Push() {
}

func (r *Repository) Checkout() {
}

func (r *Repository) CreateBranch() {
}

func (r *Repository) Reset() {
}


func (r *Repository) Filesystem() (billy.Filesystem, error) {
	w , err := r.repository.Worktree()
	if err != nil {
		return  nil, err
	}
	return w.Filesystem, nil
}
