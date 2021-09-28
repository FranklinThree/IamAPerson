package IamAPerson

import "github.com/gin-gonic/gin"

func main() {
router := gin.Default()

router.POST("/someGet", getting)
router.Run()
}


func getting(c *gin.Context){
message := c.PostForm("message")
nick := c.DefaultPostForm("nick", "anonymous")

c.JSON(200, gin.H{
"status":  "posted",
"message": message,
"nick":    nick,
})
}
