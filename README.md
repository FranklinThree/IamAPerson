#IamAPerson


主要文件说明：
    
    main.go         程序入口
    server.go       客户端
    facedatabase.go 数据存储结构
    example.go      人实例
    chioce.go       选项
    sample.go       认人脸样例（生成选择题）
    error.go        错误检查
    config.go       导入配置文件
    
    FaceData.sql    认人脸验证的数据库初始化文件
    FaceStorage.sql 学生信息的数据库初始化文件

    Net.config      外部访问配置
    Storage.config  存储学生的数据库配置
    Check.config    用于存储认人脸验证的数据库配置

接口描述：

    POST /check/upload/example && POST /storage/upload/person
        上传单个人的完整信息(多为调试用)
        /check/upload/example   上传至认人脸验证数据库
        /storage/upload/person  上传至学生信息数据库
        form-data:
            picture ([]byte形式，即文件)
            itsName
            departmentNO
            studentNumber

    GET /check/download/sample
        取得认人脸验证的实例
        json
        {
			"error": 0
			"msg":   "success"
			"data": 
            {
				"picture": 
				"truth":   
				"A":       
				"B":       
				"C":       
				"D":       
			}
			"redirect": ""
		})

    POST /storage/upload/picture?itsName=...
    POST /storage/download/person?studentNumber=...
        上传指定学生的图片
        form-data
            picture ([]byte形式，即文件)

    GET /storage/download/person?itsName=...
    GET /storage/download/person?studentNumber=...
        取得单个学生的完整信息

        
        json
        {
			"error": 0
			"msg":   "success"
			"data": 
            {
				"UID":           
				"itsName":       
				"picture":       
				"department":    
				"studentNumber": 
			}
			"redirect": ""
		})
            
        
    