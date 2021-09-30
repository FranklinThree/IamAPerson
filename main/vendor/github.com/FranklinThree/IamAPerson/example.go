package IamAPerson

import (
	//"database/sql"
	_ "github.com/gin-gonic/gin"
)

// Example 人实例
type Example struct {
	UID           int
	picture       []byte
	itsName       string
	studentNumber int
	departmentNO  int
}

func (example *Example) toData() (data []byte) {
	return nil
}
