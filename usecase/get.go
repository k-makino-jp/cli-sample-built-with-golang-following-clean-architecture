package usecase

import (
	"fmt"

	"github.com/k-makino-jp/cli-built-with-golang-following-clean-architecture/entity"
)

// GetCmdInteractor implements Cmd interface.
type GetCmdInteractor struct {
	UserCollector UserCollector
}

type UserCollector interface {
	Get() ([]entity.User, error)
}

func (g GetCmdInteractor) Execute() error {
	users, err := g.UserCollector.Get()
	if err == nil {
		fmt.Println(users)
	}
	return err
}
