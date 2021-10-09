package IamAPerson

import "C"
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
	CheckConfig   Config
	StorageConfig Config
	NetConfig     Config
}

func (server *Server) Start() (err error) {
	router := gin.New()
	var CheckDB FaceDataBase
	var StorageDB FaceDataBase

	err = StorageDB.StartByConfig(server.StorageConfig)
	if !CheckErr(err) {
		return errors.New("Storage数据库初始化错误，请检查配置文件：" + server.StorageConfig.Path)
	}

	err = CheckDB.StartByConfig(server.CheckConfig)
	if !CheckErr(err) {
		return errors.New("Check数据库初始化错误，请检查配置文件：" + server.CheckConfig.Path)

	}
	defer func() {
		err = CheckDB.database.Close()
		CheckErr(err)
	}()

	router.POST("/check/upload/example", func(c *gin.Context) {

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
		err = CheckDB.addExample(example)
		CheckErr(err)
		c.String(http.StatusOK, "上传文件成功")
	})

	router.GET("/check/download/sample", func(c *gin.Context) {
		sample, err := CheckDB.getSample()
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

	router.POST("/storage/upload/person", func(c *gin.Context) {

		pictureFile, err := c.FormFile("picture")
		if !CheckErr(err) {
			c.String(http.StatusBadRequest, "上传格式错误,找不到 picture")
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
		err = StorageDB.addExample(example)
		CheckErr(err)
		c.String(http.StatusOK, "上传文件成功")
	})

	router.POST("/storage/upload/picture", func(c *gin.Context) {

		pictureFile, err := c.FormFile("picture")
		if !CheckErr(err) {
			c.String(http.StatusBadRequest, "上传格式错误,找不到 picture")
			return
		}

		//上传初始化
		var readLength int
		var picture []byte
		buffer := make([]byte, 1024*1024)

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
			picture = append(buffer[:readLength])
		}

		inputStudentNumber := c.Query("studentNumber")
		inputItsName := c.Query("itsName")

		if inputItsName == "" && inputStudentNumber == "" {
			c.String(http.StatusBadRequest, "没有传输需要检索的学号或姓名")
			return
		}
		if inputItsName != "" {
			err = StorageDB.UPDATEOne("SET havePicture = 1,picture = "+string(picture), "WHERE itsName = "+inputItsName)
			if !CheckErr(err) {
				c.String(http.StatusBadRequest, "找不到名为\""+inputItsName+"\"的学生")
				return
			}
		} else if inputStudentNumber != "" {
			err = StorageDB.UPDATEOne("SET havePicture = 1,picture = "+string(picture), "WHERE studentNumber = "+inputStudentNumber)
			if !CheckErr(err) {
				c.String(http.StatusBadRequest, "找不到学号为\""+inputStudentNumber+"\"的学生")
				return
			}
		}

		c.String(http.StatusOK, "上传图片成功")
	})

	router.GET("/storage/download/person", func(c *gin.Context) {

		inputStudentNumber := c.Query("studentNumber")
		inputItsName := c.Query("itsName")
		var example Example
		if inputItsName == "" && inputStudentNumber == "" {
			c.String(http.StatusBadRequest, "没有传输需要检索的学号或姓名")
			return
		}

		//这里优先以姓名查询
		if inputItsName != "" {
			example, err = StorageDB.QUERYOne("WHERE itsName =" + inputItsName)
			CheckErr(err)
		} else if inputStudentNumber != "" {
			example, err = StorageDB.QUERYOne("WHERE studentNumber = " + inputStudentNumber)
			CheckErr(err)
		}

		//提取部门编号所对应的部门名称
		departmentName, err := StorageDB.getDepartmentName(example.departmentNO)
		if !CheckErr(err) {
			c.String(http.StatusInternalServerError, "服务器内部错误：找不到部门！")
			return
		}

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
	err = router.Run(server.NetConfig.Map["ip"] + ":" + server.NetConfig.Map["port"])
	if !CheckErr(err) {
		return errors.New("服务器启动异常！")
	}
	return
}
