package repository

import (
	"errors"

	"github.com/go-git/go-billy/v5"
)

var (
	ErrConfict = errors.New("a better message about the confilct error")
)

type  Repository interface {
	Commit()
	Add()
	Push()
	Pull()
	Reset()
	Clean()

	Filesystem() (billy.Filesystem, error)
}