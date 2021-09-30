package IamAPerson

type Choice struct {
	sentence string
	/**
	0				无
	1				名字
	2				部门
	3				UID
	4				学号
	*/
	choiceType int
	isRight    bool
}
