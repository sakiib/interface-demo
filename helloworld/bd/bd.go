package bd

import (
	"errors"
	"github.com/sakiib/interface-demo/helloworld/api"
)

const (
	COUNTRY = "BD"
	CAPITAL = "Dhaka"
)

type HelloWorld struct {
	store   map[string]string
	capital string
}

var _ api.HelloWorldInterface = &HelloWorld{}

func New() (*HelloWorld, error) {
	hw := &HelloWorld{
		store:   make(map[string]string),
		capital: CAPITAL,
	}
	return hw, nil
}

func (hw *HelloWorld) Get() (string, error) {
	capital, ok := hw.store[COUNTRY]
	if !ok {
		return "", errors.New("BD capital not found")
	}
	return capital, nil
}

func (hw *HelloWorld) Set(capital string) error {
	if _, ok := hw.store[COUNTRY]; ok {
		return errors.New("already exists")
	}

	hw.store[COUNTRY] = capital
	return nil
}

func (hw *HelloWorld) Capital() string {
	return hw.capital
}
