package helloworld

import (
	"errors"
	"github.com/sakiib/interface-demo/helloworld/api"
	"github.com/sakiib/interface-demo/helloworld/aus"
	"github.com/sakiib/interface-demo/helloworld/bd"
)

func NewHelloWorld(country string) (api.HelloWorldInterface, error) {
	switch country {
	case "BD":
		return bd.New()
	case "AUS":
		return aus.New()
	default:
		return nil, errors.New("unknown country code! :(")
	}
}
