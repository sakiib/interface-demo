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
	sw := &HelloWorld{
		store:   make(map[string]string),
		capital: CAPITAL,
	}
	return sw, nil
}

func (h *HelloWorld) Get() (string, error) {
	val, ok := h.store[COUNTRY]
	if !ok {
		return "", errors.New("BD capital not found")
	}
	return val, nil
}

func (h *HelloWorld) Set(capital string) error {
	if _, ok := h.store[COUNTRY]; ok {
		return errors.New("already exists")
	}

	h.store[COUNTRY] = capital
	return nil
}

func (h *HelloWorld) Capital() string {
	return h.capital
}
