package IamAPerson

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main(){
	router := gin.Default()
	router.POST("/upload",func(c *gin.Context){

		file,err := c.FormFile("upload")
		if ! CheckErr(err){
			c.String(http.StatusBadRequest,"请求失败")
			return
		}
		//获取文件名
		fileName := file.Filename
		fmt.Println("文件名：",fileName)

		//保存文件到服务器本地
		//SaveUploadedFile(文件头，保存路径)

		if err := c.SaveUploadedFile(file,fileName);! CheckErr(err){
			c.String(http.StatusBadRequest,"保存失败 Error:%s",err.Error())
			return
		}
		c.String(http.StatusOK,"上传文件成功")
	})
	router.Run()
}
//CheckErr true为正常，false为有错误
func CheckErr(err error) bool{
	if err != nil{
		panic(err)
		return false
	}
	return true
}