package IamAPerson

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

type Server struct {
}

func (server *Server) start() (err error) {
	router := gin.Default()
	var fdb FaceDataBase
	fdb.database, err = sql.Open("mysql", "root:333333@(127.0.0.1:3306)/facedata?charset=utf8")

	CheckErr(err)
	defer fdb.database.Close()
	router.POST("/upload/example", func(c *gin.Context) {

		pictureFile, err := c.FormFile("picture")
		if !CheckErr(err) {
			c.String(http.StatusBadRequest, "上传格式错误,step 1：找不到 picture")
			return
		}
		itsNameFile, err := c.FormFile("itsName")
		if !CheckErr(err) {
			c.String(http.StatusBadRequest, "上传格式错误,step 2：找不到 itsName")
			return
		}
		DepartmentNameFile, err := c.FormFile("departmentName")
		if !CheckErr(err) {
			c.String(http.StatusBadRequest, "上传格式错误,step 3：找不到 departmentName")
			return
		}
		studentNumberFile, err := c.FormFile("studentNumber")
		if !CheckErr(err) {
			c.String(http.StatusBadRequest, "上传格式错误,step 4：找不到 studentNumber")
			return
		}

		pictureFileHeader, err := pictureFile.Open()
		CheckErr(err)
		itsNameFileHeader, err := itsNameFile.Open()
		CheckErr(err)
		DepartmentNameFileHeader, err := DepartmentNameFile.Open()
		CheckErr(err)
		studentNumberFileHeader, err := studentNumberFile.Open()
		CheckErr(err)

		defer func() {
			err = pictureFileHeader.Close()
			CheckErr(err)
			err = itsNameFileHeader.Close()
			CheckErr(err)
			err = DepartmentNameFileHeader.Close()
			CheckErr(err)
			err = studentNumberFileHeader.Close()
			CheckErr(err)
		}()

		//上传初始化
		var readLength int
		var example Example
		buffer := make([]byte, 1024*1024)

		//读取图片数据
		for {
			readLength, _ = pictureFileHeader.Read(buffer)
			if readLength == 0 {
				break
			}
			example.picture = append(buffer[:readLength])
		}

		//读取名字
		readLength, err = itsNameFileHeader.Read(buffer)
		example.itsName = string(buffer[:readLength])

		//读取部门名并转化为ID
		readLength, err = DepartmentNameFileHeader.Read(buffer)
		departmentName := string(buffer[:readLength])
		example.departmentNO, err = fdb.getDepartmentNO(departmentName)

		//读取学号
		readLength, err = studentNumberFileHeader.Read(buffer)
		example.studentNumber, err = strconv.Atoi(string(buffer))

		//储存实例
		err = fdb.addExample(example)
		CheckErr(err)
		c.String(http.StatusOK, "上传文件成功")
	})

	router.GET("/download/sample", func(c *gin.Context) {

	})
	router.GET("/download/person", func(c *gin.Context) {

		//
		example, err := fdb.getExample(1)
		CheckErr(err)
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
	})
	router.Run()
	return err
}
