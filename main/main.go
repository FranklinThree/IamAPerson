package main

import (
	"fmt"
	"github.com/FranklinThree/IamAPerson"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	CheckConfig, err := IamAPerson.NewConfig("Check.config", "")
	if !IamAPerson.CheckErr(err) {

	}
	StorageConfig, err := IamAPerson.NewConfig("Storage.config", "Storage")
	NetConfig, err := NewConfig("IPConfig", "Net")
	gin.SetMode(gin.DebugMode)
	server := Server{CheckConfig, StorageConfig, NetConfig}
	err = server.Start()
	fmt.Println("result = ", err)
}
