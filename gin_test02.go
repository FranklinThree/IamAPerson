package IamAPerson

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	router := gin.Default()

	router.GET("/someGet", getting)
	router.Run()
}
func getting(c *gin.Context){
	c.String(http.StatusOK,"Hello gin")
}


