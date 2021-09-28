package IamAPerson

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
router := gin.Default()

router.GET("/someGet/:name", getting)
router.Run()
}


func getting(c *gin.Context){
name:=c.Param("name")
c.String(http.StatusOK,"Hello %s",name)
}

