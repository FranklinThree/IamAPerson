package IamAPerson

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

// FaceDataBase 人脸实例数据库
type FaceDataBase struct {
	database           *sql.DB
	PersonCapacity     int
	DepartmentCapacity int
}

// Start 初始化
func (fdb *FaceDataBase) Start(driverName string, dataSourceName string) (err error) {

	//初始化数据库
	fdb.database, err = sql.Open(driverName, dataSourceName)
	if !CheckErr(err) {
		return errors.New("数据库初始化失败")
	}

	err = fdb.maintain()
	if !CheckErr(err) {
		return errors.New("数据库总数据抽取失败")
	}

	fmt.Println("数据库初始化成功")

	return err
}
func (fdb *FaceDataBase) StartByConfig(config Config) (err error) {

	//初始化数据库
	fdb.database, err = sql.Open(config.Map["driverName"],
		config.Map["userName"]+":"+config.Map["userKey"]+
			"@("+config.Map["ip"]+":"+config.Map["port"]+")/"+
			config.Map["databaseName"]+
			"?charset="+config.Map["charset"],
	)

	if !CheckErr(err) {
		return errors.New("数据库初始化失败")
	}

	err = fdb.maintain()
	if !CheckErr(err) {
		return errors.New("数据库总数据抽取失败")
	}

	fmt.Println("数据库初始化成功")

	return err
}

// maintain 维护
func (fdb *FaceDataBase) maintain() (err error) {

	//统计人脸实例规模
	personCountQuery, err := fdb.database.Query("SELECT count(UID) FROM person")
	CheckErr(err)

	personCountQuery.Next()
	err = personCountQuery.Scan(&fdb.PersonCapacity)
	CheckErr(err)

	fmt.Println("数据库人实例总数抽取成功：" + strconv.Itoa(fdb.PersonCapacity))
	//统计部门规模
	DepartmentCountQuery, err := fdb.database.Query("SELECT count(departmentNO) FROM department")
	CheckErr(err)

	DepartmentCountQuery.Next()
	err = DepartmentCountQuery.Scan(&fdb.DepartmentCapacity)
	CheckErr(err)

	fmt.Println("数据库部门总数抽取成功：" + strconv.Itoa(fdb.DepartmentCapacity))
	defer func() {
		err = personCountQuery.Close()
		CheckErr(err)
		err = DepartmentCountQuery.Close()
		CheckErr(err)
	}()
	return err
}

// addExample 添加实例
func (fdb *FaceDataBase) addExample(example Example) (err error) {

	personInsert, err := fdb.database.Prepare("INSERT person SET itsName=?,departmentNO=?,picture=?,havePicture=?,studentNumber=?")
	CheckErr(err)

	_, err = personInsert.Exec(example.itsName, example.departmentNO, example.picture, example.picture != nil, example.studentNumber)
	if !CheckErr(err) {
		return errors.New("数据插入错误！")
	}

	fdb.PersonCapacity++
	return nil
}

// getExample 取得指定实例
func (fdb *FaceDataBase) getExample(UID int, havePicture bool) (example Example, err error) {

	example.UID = UID
	if havePicture {

		//查询人脸数据表picture
		personQuery, err := fdb.database.Query("SELECT itsName,picture,departmentNO,studentNumber FROM person WHERE UID=" + strconv.Itoa(UID))
		CheckErr(err)

		defer func() {
			err = personQuery.Close()
			CheckErr(err)
		}()

		//查找 UID 所对应的人脸实例
		personQuery.Next()
		err = personQuery.Scan(&example.itsName, &example.picture, &example.departmentNO, &example.studentNumber)
		CheckErr(err)

	} else {
		//查询人脸数据表picture
		personQuery, err := fdb.database.Query("SELECT itsName,departmentNO,studentNumber FROM person WHERE UID=" + strconv.Itoa(UID))
		CheckErr(err)

		defer func() {
			err = personQuery.Close()
			CheckErr(err)
		}()

		//查找 UID 所对应的人脸实例
		personQuery.Next()
		err = personQuery.Scan(&example.itsName, &example.departmentNO, &example.studentNumber)
		CheckErr(err)

	}
	return
}

// getRandomExample 取得随机实例
func (fdb *FaceDataBase) getRandomExample(getPicture bool, WHEREExpression string) (example Example, err error) {

	NO := rand.Intn(fdb.PersonCapacity)
	if getPicture {

		//查询人脸数据表picture
		personQuery, err := fdb.database.Query("SELECT itsName,picture,departmentNO,studentNumber FROM person " + WHEREExpression)
		CheckErr(err)

		defer func() {
			err = personQuery.Close()
			CheckErr(err)
		}()

		//查找 顺序NO 所对应的人脸实例
		for i := 0; i < NO; i++ {
			personQuery.Next()
		}
		err = personQuery.Scan(&example.itsName, &example.picture, &example.departmentNO, &example.studentNumber)
		CheckErr(err)

	} else {
		//查询人脸数据表picture
		personQuery, err := fdb.database.Query("SELECT itsName,departmentNO,studentNumber FROM person " + WHEREExpression)
		CheckErr(err)

		defer func() {
			err = personQuery.Close()
			CheckErr(err)
		}()

		//查找 UID 所对应的人脸实例
		personQuery.Next()
		err = personQuery.Scan(&example.itsName, &example.departmentNO, &example.studentNumber)
		CheckErr(err)

	}
	return
}

// getUID 找到名字所对应的UID
func (fdb *FaceDataBase) getUID(itsName string) (UID int, err error) {
	PersonQuery, err := fdb.database.Query("SELECT UID FROM person WHERE itsName LIKE " + itsName)
	UID = 0
	for PersonQuery.Next() {
		if UID != 0 {
			return 0, errors.New("有多个数据与该姓名匹配：" + itsName)
		}
		err = PersonQuery.Scan(&UID)
		CheckErr(err)
	}
	return UID, err
}

// addDepartment 添加一个部门
func (fdb *FaceDataBase) addDepartment(departmentName string) (res sql.Result, err error) {

	DepartmentInsert, err := fdb.database.Prepare("INSERT department SET departmentName=?")
	if !CheckErr(err) {
		return nil, errors.New("未找到department数据表")
	}
	res, err = DepartmentInsert.Exec(departmentName)
	if !CheckErr(err) {
		return nil, errors.New("无法创建部门：" + departmentName)
	}
	fdb.DepartmentCapacity++
	return

}

// checkDepartment 检查部门是否存在
func (fdb *FaceDataBase) checkDepartment(departmentNO int) (err error) {

	if departmentNO < 1 {
		return errors.New("部门编号不合法：" + strconv.Itoa(departmentNO))
	}

	DepartmentQuery, err := fdb.database.Query("SELECT departmentNO from department")
	for DepartmentQuery.Next() {
		var departmentNOTemp int
		err = DepartmentQuery.Scan(&departmentNOTemp)
		CheckErr(err)
		if departmentNO == departmentNOTemp {
			return nil
		}
	}
	return errors.New("未找到对应的部门:" + strconv.Itoa(departmentNO))
}

//	checkHavePicture 检查此人是否有图片
func (fdb *FaceDataBase) checkHavePicture(UID int) (have bool, err error) {
	PersonQuery, err := fdb.database.Query("SELECT havePicture FROM person WHERE UID=" + strconv.Itoa(UID))
	defer func() {
		err = PersonQuery.Close()
		CheckErr(err)
	}()
	CheckErr(err)
	PersonQuery.Next()
	err = PersonQuery.Scan(&have)
	CheckErr(err)
	return
}

// getDepartmentName 从部门编号转换为部门名称
func (fdb *FaceDataBase) getDepartmentName(departmentNO int) (departmentName string, err error) {

	DepartmentQuery, err := fdb.database.Query("SELECT departmentName FROM department WHERE departmentNO =" + strconv.Itoa(departmentNO))
	CheckErr(err)

	defer func() {
		err = DepartmentQuery.Close()
		CheckErr(err)
	}()
	for DepartmentQuery.Next() {
		err = DepartmentQuery.Scan(&departmentName)
		CheckErr(err)
	}

	return
}

// getDepartmentNO 从部门名称转换为部门编号
func (fdb *FaceDataBase) getDepartmentNO(departmentName string) (departmentNO int, err error) {
	DepartmentQuery, err := fdb.database.Query("SELECT departmentID FROM department " +
		"WHERE (departmentName LIKE %" + departmentName + "%) OR (" + departmentName + " LIKE %departmentName%)")
	if !CheckErr(err) {
		return 0, errors.New("department数据库初始化失败：")
	}
	departmentNO = 0
	for DepartmentQuery.Next() {
		if departmentNO != 0 {
			return -1, errors.New("输入的部门名称错误：")
		}
		err = DepartmentQuery.Scan(&departmentNO)
		if !CheckErr(err) {
			return -2, errors.New("这里怎么会有错误呢")
		}
	}
	return departmentNO, nil
}

// getSample 取得一个测试单元
func (fdb *FaceDataBase) getSample() (sample Sample, err error) {

	if fdb.PersonCapacity < 4 {
		return Sample{}, errors.New("数据库没有足够的人实例：" + strconv.Itoa(fdb.PersonCapacity))
	}
	if fdb.DepartmentCapacity < 2 {
		return Sample{}, errors.New("数据库没有足够的部门：" + strconv.Itoa(fdb.DepartmentCapacity))
	}
	var choices [4]Choice
	var examples [4]Example

	//取得图片的实例
	//不安全的一个随机数
	example, err := fdb.getRandomExample(true, "")
	trueDepartmentName, err := fdb.getDepartmentName(example.departmentNO)
	CheckErr(err)

	//fmt.Println("Example:",example)

	//先定义四个错误选项
	departmentNameTemp, err := fdb.getDepartmentName((example.departmentNO+rand.Intn(fdb.DepartmentCapacity)-1)%fdb.DepartmentCapacity + 1)
	choices[0] = Choice{
		departmentNameTemp,
		DepartmentNO,
		false,
	}
	for i := 0; i < 3; i++ {
		examples[i], err = fdb.getRandomExample(false, "WHERE UID <> "+strconv.Itoa(example.UID))
		if examples[i].UID == example.UID {
			i--
			continue
		}

	}
	choices[1] = Choice{
		examples[0].itsName,
		ItsName,
		false,
	}
	choices[2] = Choice{
		strconv.Itoa(examples[1].UID),
		UID,
		false,
	}
	choices[3] = Choice{
		strconv.Itoa(examples[2].studentNumber),
		StudentNumber,
		false,
	}

	//随机替换一个选项为正确选项
	randomTrue := rand.Intn(4)
	switch randomTrue {
	case 0:
		choices[0].sentence = trueDepartmentName
		choices[0].isRight = true
	case 1:
		choices[1].sentence = example.itsName
		choices[1].isRight = true
	case 2:
		choices[2].sentence = strconv.Itoa(example.UID)
		choices[2].isRight = true
	case 3:
		choices[3].sentence = strconv.Itoa(example.studentNumber)
		choices[3].isRight = true
	default:
		return sample, errors.New("这怎么可能呢。")
	}

	//随机打乱这些选项
	var mixer = [4]int{0, 1, 2, 3}
	for i := 0; i < 16; i++ {
		a := rand.Intn(4-1) + 1
		b := rand.Intn(4-1) + 1
		mixer[a], mixer[b] = mixer[b], mixer[a]
	}
	for i := 0; i < 4; i++ {
		sample.choices[i] = choices[mixer[i]]
	}

	//找到正确选项
	err = sample.getTheTrue()
	CheckErr(err)

	//我很纠结
	sample.picture = &example.picture

	return sample, nil
}

//以下函数仅供person表操作单个example用

// DELETEOne 删除实例
func (fdb *FaceDataBase) DELETEOne(WHEREExpression string) (err error) {
	PersonDelete, err := fdb.database.Prepare("DELETE FROM person " + WHEREExpression)
	_, err = PersonDelete.Exec()
	fdb.PersonCapacity--
	return err
}

// UPDATEOne 更新实例
func (fdb *FaceDataBase) UPDATEOne(SETExpression string, WHEREExpression string) (err error) {
	PersonUpdate, err := fdb.database.Prepare("UPDATE person " + SETExpression + " " + WHEREExpression)
	_, err = PersonUpdate.Exec()
	CheckErr(err)
	return
}

// QUERYOne 查找实例
func (fdb *FaceDataBase) QUERYOne(WHEREExpression string) (example Example, err error) {
	QueryPerson, err := fdb.database.Query("SELECT UID,picture,departmentNO,studentNumber,itsName FROM person " + WHEREExpression)
	defer func() {
		err = QueryPerson.Close()
		CheckErr(err)
	}()
	QueryPerson.Next()
	err = QueryPerson.Scan(&example.UID, &example.picture, &example.departmentNO, &example.studentNumber, &example.itsName)
	CheckErr(err)

	return
}

// QUERYOneNoPic 查找实例，无需传输图片
func (fdb *FaceDataBase) QUERYOneNoPic(WHEREExpression string) (example Example, err error) {
	QueryPerson, err := fdb.database.Query("SELECT UID,departmentNO,studentNumber,itsName FROM person " + WHEREExpression)
	defer func() {
		err = QueryPerson.Close()
		CheckErr(err)
	}()
	QueryPerson.Next()
	err = QueryPerson.Scan(&example.UID, &example.departmentNO, &example.studentNumber, &example.itsName)
	CheckErr(err)

	return
}
