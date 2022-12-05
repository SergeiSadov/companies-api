package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type nonInterfaceCommand struct {
}

type TestCommand struct {
	Ctx context.Context
}

func (cmd *TestCommand) Exec(client *gorm.DB) (err error) {
	return nil
}

func TestInteractor_Wire_Success(t *testing.T) {
	interactor := NewInteractor(&gorm.DB{})
	assert.NoError(t, interactor.Perform(&TestCommand{}))
}

func TestInteractor_Wire_Failed(t *testing.T) {
	interactor := NewInteractor(&gorm.DB{})
	assert.Error(t, interactor.Perform(&nonInterfaceCommand{}))
}
