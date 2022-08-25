package service

import (
	"encoding/json"
	"fmt"
	"github.com/shawu21/test/model"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var loginUrl = "https://dev-auth.itoken.team/Auth/Login"
var infoUrl = "https://dev-auth.itoken.team/Profile"

//func LoginService(userLogin *model.UserLogin) (string, error) {
//	buf := &bytes.Buffer{}
//	bodywrite := multipart.NewWriter(buf)
//	bodywrite.WriteField("email", userLogin.Email)
//	bodywrite.WriteField("secret", userLogin.Password)
//	contentType := bodywrite.FormDataContentType()
//	err := bodywrite.Close() //不能用defer，在请求体完成之后，需要将结尾符补上
//	if err != nil {
//		return nil, err
//	}
//	cc := &http.Client{}
//	req, _ := http.NewRequest(http.MethodPost, loginUrl, buf)
//	req.Header.Set("Content-Type", contentType)
//	resp, err := cc.Do(req)
//	defer resp.Body.Close()
//	return resp, err
//}

func GetInfo(token string) (*model.User, error) {
	var user *model.User
	req, _ := http.NewRequest(http.MethodGet, infoUrl, nil)
	accessToken := fmt.Sprintf("Bearer %s", token)
	req.Header.Add("Authorization", accessToken)
	cc := &http.Client{}
	resp, err := cc.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	res := make(map[string]interface{})
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	user.Name = res["realName"].(string)
	user.Email = res["email"].(string)
	user.Tel = res["phone"].(string)
	user.QQ = res["qq"].(string)
	user.Gender = res["gender"].(int)
	user.IdcardNumber = res["studentNumber"].(string)
	return user, nil
}

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
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(body, &res)
	token := res["accessToken"].(string)
	return token, err
}
