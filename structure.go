package IamAPerson

import (
	"database/sql"
)

// Example 人脸实例
type Example struct {
	UID int
	picture []byte
	iname string
	department string
}

func (example *Example)toData() (data []byte){
	return nil
}

// FaceDataBase 人脸实例数据库
type FaceDataBase struct {
	database *sql.DB
	FaceCapacity int
	DepartmentCapacity int
}

// Start 人脸实例集合：初始化
func (fdb FaceDataBase) Start(driverName string,dataSourceName string) error{
	var err error

	//初始化数据库
	fdb.database,err = sql.Open(driverName,dataSourceName)
	CheckErr(err)

	//统计人脸实例规模
	PictureCountQuery,err := fdb.database.Query("SELECT count(UID) FROM picture")
	CheckErr(err)
	for PictureCountQuery.Next(){
		fdb.FaceCapacity++
	}

	//统计部门规模
	DepartmentCountQuery,err := fdb.database.Query("SELECT count(departmentNO) FROM department")
	CheckErr(err)
	for DepartmentCountQuery.Next(){
		fdb.DepartmentCapacity++
	}
	defer func(){
		err = PictureCountQuery.Close()
		CheckErr(err)
		err = DepartmentCountQuery.Close()
		CheckErr(err)
	}()
	return err
}

//Maintain 人脸实例集合：维护
func (fdb FaceDataBase) Maintain() error{
	var err error
	//统计人脸元素规模
	PictureCountQuery,err := fdb.database.Query("SELECT count(UID) FROM picture")
	CheckErr(err)
	for PictureCountQuery.Next(){
		fdb.FaceCapacity++
	}

	//统计部门规模
	DepartmentCountQuery,err := fdb.database.Query("SELECT count(departmentNO) FROM department")
	CheckErr(err)
	for DepartmentCountQuery.Next(){
		fdb.DepartmentCapacity++
	}
	defer func(){
		err = PictureCountQuery.Close()
		CheckErr(err)
		err = DepartmentCountQuery.Close()
		CheckErr(err)
	}()
	return err
}
// AddExample 人脸实例数据库：添加实例
func (fdb FaceDataBase) AddExample (example Example)  {

}

// GetExample 人脸实例数据库：取得指定实例
func (fdb FaceDataBase) GetExample(UID int) (example Example,err error){
	//查询人脸数据表picture
	PictureQuery,err := fdb.database.Query("SELECT iname,picture,departmentNO FROM picture WHERE UID=?")
	CheckErr(err)
	//查询部门数据表department
	DepartmentQuery,err := fdb.database.Query("SELECT departmentNO FROM department WHERE UID=?")
	CheckErr(err)

	defer func(){
		err = PictureQuery.Close()
		CheckErr(err)
		err = DepartmentQuery.Close()
		CheckErr(err)
	}()

	//查找 UID 所对应的人脸实例
	for PictureQuery.Next(){
		var departmentNO int
		err = PictureQuery.Scan(1,&example.iname,&example.picture,&departmentNO)
		CheckErr(err)
		for DepartmentQuery.Next(){
			DepartmentQuery.Scan(departmentNO,&example.department)
			CheckErr(err)
		}
	}
	return example,err
}