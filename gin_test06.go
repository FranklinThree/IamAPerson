package IamAPerson


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
router := gin.Default()

router.POST("/someGet", getting)
router.Run()
}


func getting(c *gin.Context){

id := c.Query("id")
page := c.DefaultQuery("page", "0")
name := c.PostForm("name")
message := c.PostForm("message")

c.String(http.StatusOK,"id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}

