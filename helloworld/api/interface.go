package api

type HelloWorldInterface interface {
	Get() (string, error)
	Set(string) error
	Capital() string
}
