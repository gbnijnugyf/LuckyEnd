package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/shawu21/LuckyBackend/common"
	"github.com/shawu21/LuckyBackend/model"
	log "github.com/sirupsen/logrus"
)

var loginUrl = "https://dev-auth.itoken.team/Auth/Login"
var infoUrl = "https://dev-auth.itoken.team/Profile"

func GetInfo(token string) (*model.User, error) {
	user := &model.User{}
	req, _ := http.NewRequest(http.MethodGet, infoUrl, nil)
	accessToken := fmt.Sprintf("Bearer %s", token)
	req.Header.Add("Authorization", accessToken)
	cc := &http.Client{}
	resp, err := cc.Do(req)
	if err != nil {
		log.Errorf("client do error %+v", err)
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("io read error %+v", err)
		return nil, err
	}
	defer resp.Body.Close()
	res := make(map[string]interface{})
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Errorf("unmarshal error %+v", err)
		return nil, err
	}
	user.Name = res["realName"].(string)
	user.Email = res["email"].(string)
	user.Tel = res["phone"].(string)
	user.QQ = res["qq"].(string)
	user.Gender = int(res["gender"].(float64)) // interface无法直接转为int类型
	user.IdcardNumber = res["id"].(string)
	user.School = common.Whut
	return user, nil
}

func SendForm(email, secret string) (string, error) {
	data := url.Values{}
	data.Add("email", email)
	data.Add("secret", secret)
	req, err := http.NewRequest(http.MethodPost, loginUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Errorf("request error %+v", errors.WithStack(err))
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	cc := http.Client{}
	resp, err := cc.Do(req)
	if err != nil {
		log.Errorf("client do error %+v", err)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Errorf("io read error %+v", err)
		return "", err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(body, &res)
	token, ok := res["accessToken"].(string)
	if !ok {
		log.Errorf("Invalid accesstoken err %+v", errors.WithStack(err))
		return "", err
	}
	return token, err
}
