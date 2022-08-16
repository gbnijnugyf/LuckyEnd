package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"test/common"
	"test/helper"
	"test/model"

	"github.com/gin-gonic/gin"
)

type login struct {
	Email    string
	Password string
}

func Login(c *gin.Context) {
	var userLogin login
	var User *model.User
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", nil))
		return
	}
	loginUrl := "https://auth.itoken.team/Auth/Login"
	buf := &bytes.Buffer{}
	bodywrite := multipart.NewWriter(buf)
	bodywrite.WriteField("email", userLogin.Email)
	bodywrite.WriteField("secret", userLogin.Password)
	contentType := bodywrite.FormDataContentType()
	bodywrite.Close() //不能用defer，在请求体完成之后，需要将结尾符补上
	cc := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, loginUrl, buf)
	req.Header.Set("Content-Type", contentType)
	resp, _ := cc.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.Status {
	case common.NoExist:
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "用户不存在", nil))
	case common.PasError:
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "密码错误", nil))
	case common.LoginSuccess:
		// 判断用户是否已经注册
		if err := model.UserCheck(userLogin.Email); err != nil {
			if err := c.ShouldBindJSON(User); err != nil {
				c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定错误", nil))
			}

		}
	}
	fmt.Println(resp)
	fmt.Println(string(body))
}

// func CheckUserEmail(c *gin.Context) {
// 	UserID := c.MustGet("user_id").(int)
// 	UserEmail := model.GetUserEmailByUserID(UserID)
// 	if UserEmail == "" {
// 		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "未绑定邮箱", nil))
// 		return
// 	}
// 	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "邮箱已绑定", UserEmail))
// }

// func BindEmail(c *gin.Context) {
// 	var user model.User
// 	student_number := c.MustGet("student_number").(string)
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据绑定失败", err.Error()))
// 		return
// 	}
// 	user.IdcardNumber = student_number
// 	if res := model.BindEmail(user); res.Status == common.CodeError {
// 		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "邮箱绑定失败", res.Data))
// 		return
// 	}
// 	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "绑定邮箱成功", user.Email))
// }
