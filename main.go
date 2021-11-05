package main

import (
	"fmt"
	hello "github.com/sakiib/interface-demo/helloworld"
)

func main() {
	countries := []string{"BD", "AUS", "BAD"}
	for _, country := range countries {
		countryInterface, err := hello.NewHelloWorld(country)
		if err != nil {
			fmt.Println("error: ", err.Error())
			continue
		}

		capital, err := countryInterface.Get()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("cap: ", capital)
		}

		if err := countryInterface.Set(countryInterface.Capital()); err != nil {
			fmt.Println("error: ", err.Error())
			continue
		}

		capital, err = countryInterface.Get()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("cap: ", capital)
		}
	}
}
