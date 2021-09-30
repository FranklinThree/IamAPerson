package IamAPerson

const (
	ItsName       = 1
	DepartmentNO  = 2
	UID           = 3
	StudentNumber = 4
)

func TypetoString(t int) (result string) {
	switch t {
	case 1:
		return "itsName"
	case 2:
		return "departmentNO"
	case 3:
		return "UID"
	case 4:
		return "studentNumber"
	}
	return
}
func TypetoInt(t string) (result int) {
	switch t {
	case "itsName":
		return 1
	case "departmentNO":
		return 2
	case "UID":
		return 3
	case "studentNumber":
		return 4
	}
	return
}
