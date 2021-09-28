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
firstname := c.DefaultQuery("firstname", "Guest")
lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

