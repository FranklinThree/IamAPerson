package IamAPerson

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"

	"net/http"
)
func main(){
	router := gin.Default()
	var fdb FaceDataBase
	var err error
	fdb.database,err = sql.Open("mysql","root:333333@(127.0.0.1:3306)/facedata?charset=utf8")

	CheckErr(err)
	defer fdb.database.Close()
	//router.POST("/upload",func(c *gin.Context){
	//
	//	file,err := c.FormFile("upload")
	//	//iname,err := c.FormFile("i")
	//	if ! CheckErr(err){
	//		c.String(http.StatusBadRequest,"请求失败")
	//		return
	//	}
	//	reader,err := file.Open()
	//	if ! CheckErr(err){
	//		c.String(http.StatusBadRequest,"文件写入初始化失败")
	//		return
	//	}
	//	buffer := make([]byte,1024*1024)
	//	updates,err := db.Prepare("UPDATE picture SET picture = ? WHERE UID = ?")
	//	CheckErr(err)
	//	res,err := updates.Exec("",1)
	//	CheckErr(err)
	//	fmt.Println(res)
	//	//writer 仅供测试
	//	//writer,_ := os.Create("1.jpg")
	//	for {
	//		readlength,_ :=reader.Read(buffer)
	//		if readlength == 0 {
	//			break
	//		}
	//
	//		updates,err := db.Prepare("UPDATE picture SET picture = CONCAT(picture,?) WHERE UID = ?")
	//		CheckErr(err)
	//		res,err := updates.Exec(string(buffer[:readlength]),1)
	//		CheckErr(err)
	//		fmt.Println("POST res =",res)
	//		//writer.Write(buffer[:readlength])
	//	}
	//	reader.Close()
	//	//writer.Close()
	//	c.String(http.StatusOK,"上传文件成功")
	//})

	router.GET("/Download",func(c *gin.Context){

		example,err := fdb.GetExample(rand.Intn(fdb.FaceCapacity-1)+1)
		CheckErr(err)

		c.JSON(http.StatusOK,gin.H{
			"error":0,
			"msg":"success",
			"data":example.toData(),
			"redirect":"",
		})

	})
	router.Run()
}
