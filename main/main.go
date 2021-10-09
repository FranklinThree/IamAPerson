package main

import (
	"fmt"
	"github.com/FranklinThree/IamAPerson"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	CheckConfig, err := IamAPerson.NewConfig("Check.config", "Check")
	if !IamAPerson.CheckErr(err) {

	}
	StorageConfig, err := IamAPerson.NewConfig("Storage.config", "Storage")
	NetConfig, err := IamAPerson.NewConfig("Net.Config", "Net")
	gin.SetMode(gin.DebugMode)
	server := IamAPerson.Server{CheckConfig: CheckConfig, StorageConfig: StorageConfig, NetConfig: NetConfig}
	err = server.Start()
	fmt.Println("result = ", err)
}
