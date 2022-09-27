package db

type ICommand interface{}

type IClient interface {
	Perform(ICommand) error
}
