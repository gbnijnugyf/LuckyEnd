package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
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
	loginUrl := "https://dev-auth.itoken.team/Auth/Login"
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)
	defer writer.Close()
	emailValue, err := writer.CreateFormField("email")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	secretValue, _ := writer.CreateFormField("secret")
	_, _ = emailValue.Write([]byte("2982271907@qq.com"))
	_, _ = secretValue.Write([]byte("lxt237156"))
	req, err := http.NewRequest(http.MethodPost, loginUrl, reqBody)
	if err != nil {
		fmt.Println("login error:" + err.Error())
		return
	}
	cc := http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := cc.Do(req)
	if err != nil {
		fmt.Println("client error:" + err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error:" + err.Error())
	}
	fmt.Println(string(body))
}

func Login2() {
	loginUrl := "https://dev-auth.itoken.team/Auth/Login"
	buf := &bytes.Buffer{}
	bodywrite := multipart.NewWriter(buf)
	bodywrite.WriteField("email", "2982271907@qq.com")
	bodywrite.WriteField("secret", "lxt237156")
	contentType := bodywrite.FormDataContentType()
	bodywrite.Close() //不能用defer，在请求体完成之后，需要将结尾符补上
	cc := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, loginUrl, buf)
	req.Header.Set("Content-Type", contentType)
	resp, _ := cc.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}

func main() {
	Login2()
}
