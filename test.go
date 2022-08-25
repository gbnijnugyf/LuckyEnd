package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

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
	fmt.Println(resp.StatusCode)
	res := make(map[string]interface{})
	err := json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res["accessToken"])
	fmt.Println(string(body))
}

func GetInfo(token string) {
	InfoUrl := "https://dev-auth.itoken.team/Profile"
	req, _ := http.NewRequest(http.MethodGet, InfoUrl, nil)
	accessToken := fmt.Sprintf("Bearer %s", token)
	req.Header.Add("Authorization", accessToken)
	cc := &http.Client{}
	resp, err := cc.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	res := make(map[string]interface{})
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

var loginUrl string = "https://dev-auth.itoken.team/Auth/Login"

func SendForm(email, secret string) (string, error) {
	data := url.Values{}

	data.Add("email", email)
	data.Add("secret", secret)
	req, err := http.NewRequest(http.MethodPost, loginUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	cc := http.Client{}
	resp, err := cc.Do(req)
	if err != nil {
		return "", err
	}
	body := make([]byte, 0)
	_, err = resp.Body.Read(body)
	if err != nil {
		return "", err
	}
	// 再处理body
	res := make(map[string]interface{})
	err = json.Unmarshal(body, &res)
	fmt.Println(res["accessToken"])
	token := res["accessToken"].(string)
	return token, err
}

func main() {
	//Login2()
	//GetInfo("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjI5ODIyNzE5MDdAcXEuY29tIiwiQXZhdGFyVXJsIjoiIiwiTmljayI6IiIsIlRva2VuVHlwZSI6IkFjY2Vzc1Rva2VuIiwiaHR0cDovL3NjaGVtYXMueG1sc29hcC5vcmcvd3MvMjAwNS8wNS9pZGVudGl0eS9jbGFpbXMvbmFtZWlkZW50aWZpZXIiOiIwOGRhN2U4Ni05Nzk2LTQ1ZTAtODFhMi1iM2EyY2MzNjQ2OGIiLCJuYmYiOjE2NjEzOTgwNDMsImV4cCI6MTY2MTM5ODY0MywiaXNzIjoiaXd1dC1iYWNrZW5kIiwiYXVkIjoiaXd1dC1hcHAifQ.OhV03qSYvU9k3tq_D0usR1Tc_Dr1_p02XJlmVms-_Pg")
	token, err := SendForm("2982271907@qq.com", "lxt237156")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
}
