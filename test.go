package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// type login struct {
// 	Email    string
// 	Password string
// }

func Login() {
	// if err := c.ShouldBindJSON(&userLogin); err != nil {
	// 	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", nil))
	// 	return
	// }
	loginUrl := "https://auth.itoken.team/Auth/Login"
	payload := url.Values{"email": {"2982271907@qq.com"}, "secret": {"237156"}}
	res, err := http.PostForm(loginUrl, payload)
	if err != nil {
		fmt.Println("login error:" + err.Error())
		return
	}
	fmt.Println(res)
}

func main() {
	Login()
}
