package db

import (
	"errors"

	"gorm.io/gorm"
)

type ICmd interface {
	Exec(client *gorm.DB) (err error)
}

type Interactor struct {
	db *gorm.DB
}

func NewInteractor(db *gorm.DB) *Interactor {
	return &Interactor{db: db}
}

func (i *Interactor) Perform(cmd ICommand) (err error) {
	if command, ok := cmd.(ICmd); ok {
		return command.Exec(i.db)
	}

	return errors.New("cannot execute non gorm command")
}
