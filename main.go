package main

import (
	"fmt"
	"github.com/FranklinThree/IamAPerson"
)

func main() {
	server := IamAPerson.Server{}
	err := server.Start()
	fmt.Println("result = ", err)
}
