package controller

import (
	"test/common"
	"test/model"
)

func GeuUserIDFromDB(student_number string) int {
	res := model.GetUserIDByStudentNumber(student_number)
	if res.Status == common.CodeError {
		return common.CodeError
	}
	return res.Data.(model.User).ID
}
