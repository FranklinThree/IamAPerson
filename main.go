package IamAPerson

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	CheckConfig,err := NewConfig("Check.config","")
	if !CheckErr(err){

	}
	StorageConfig,err := NewConfig("Storage.config","Storage")
	NetConfig,err := NewConfig("IPConfig","Net")
	gin.SetMode(gin.DebugMode)
	server := Server{CheckConfig,StorageConfig,NetConfig}
	err = server.Start()
	fmt.Println("result = ", err)
}

