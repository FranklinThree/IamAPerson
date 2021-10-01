package IamAPerson

import (
	//"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	//"fmt"
)

type Server struct {
}

func (server *Server) Start() (err error) {
	router := gin.New()
	//err = http.ListenAndServe("127.0.0.1:8080", router)
	//if !CheckErr(err) {
	//	return errors.New("服务器初始化出错")
	//}
	var fdb FaceDataBase
	err = fdb.Start("mysql", "root:333333@(127.0.0.1:3306)/facedata?charset=utf8")
	CheckErr(err)
	defer fdb.database.Close()
	router.POST("/upload/example", func(c *gin.Context) {

		pictureFile, err := c.FormFile("picture")
		if !CheckErr(err) {
			c.String(http.StatusBadRequest, "上传格式错误,step 1：找不到 picture")
			return
		}

		//上传初始化
		var readLength int
		var example Example
		buffer := make([]byte, 1024*1024)

		example.itsName, _ = c.GetPostForm("itsName")

		departmentNO, _ := c.GetPostForm("departmentNO")
		example.departmentNO, err = strconv.Atoi(departmentNO)
		CheckErr(err)

		studentNumber, _ := c.GetPostForm("studentNumber")
		example.studentNumber, err = strconv.Atoi(studentNumber)
		CheckErr(err)

		pictureFileHeader, err := pictureFile.Open()
		CheckErr(err)

		defer func() {
			err = pictureFileHeader.Close()
			CheckErr(err)
		}()

		//读取图片数据
		for {
			readLength, _ = pictureFileHeader.Read(buffer)
			if readLength == 0 {
				break
			}
			example.picture = append(buffer[:readLength])
		}

		//储存实例
		err = fdb.addExample(example)
		CheckErr(err)
		c.String(http.StatusOK, "上传文件成功")
	})

	router.GET("/download/sample", func(c *gin.Context) {
		sample, err := fdb.getSample()
		CheckErr(err)
		c.JSON(http.StatusOK, gin.H{
			"error": 0,
			"msg":   "success",
			"data": gin.H{
				"picture": sample.picture,
				"truth":   sample.theTrue,
				"A":       sample.choices[0].sentence,
				"B":       sample.choices[1].sentence,
				"C":       sample.choices[2].sentence,
				"D":       sample.choices[3].sentence,
			},
			"redirect": "",
		})
	})
	router.GET("/download/person", func(c *gin.Context) {

		uid, err := strconv.Atoi(c.Query("UID"))
		CheckErr(err)
		example, err := fdb.getExample(uid, true)
		if !CheckErr(err) {
			panic(errors.New("找不到UID目标:" + strconv.Itoa(uid)))
			return
		}
		departmentName, err := fdb.getDepartmentName(example.departmentNO)
		CheckErr(err)
		c.JSON(http.StatusOK, gin.H{
			"error": 0,
			"msg":   "success",
			"data": gin.H{
				"UID":           example.UID,
				"itsName":       example.itsName,
				"picture":       example.picture,
				"department":    departmentName,
				"studentNumber": example.studentNumber,
			},
			"redirect": "",
		})
		c.String(http.StatusOK, "拉取信息成功")
	})
	router.Run("127.0.0.1:8080")
	return err
}
