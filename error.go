package IamAPerson

//CheckErr true为正常，false为有错误
func CheckErr(err error) bool{
	if err != nil{
		panic(err)
		return false
	}
	return true
}