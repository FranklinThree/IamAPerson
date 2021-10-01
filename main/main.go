package main

import (
	"fmt"
	"github.com/FranklinThree/IamAPerson"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	server := IamAPerson.Server{}
	err := server.Start()
	fmt.Println("result = ", err)
}
