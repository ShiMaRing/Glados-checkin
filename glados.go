package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type resBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	List    []struct {
		ID       int    `json:"id"`
		UserID   int    `json:"user_id"`
		Time     int64  `json:"time"`
		Asset    string `json:"asset"`
		Business string `json:"business"`
		Change   string `json:"change"`
		Balance  string `json:"balance"`
		Detail   string `json:"detail"`
	} `json:"list"`
}

const DefaultExecTime = 20
const DefaultTimeout = 10

var req *http.Request
var cookies []*http.Cookie

func init() {
	err := configInit()
	if err != nil {
		panic(err)
	}
}

func checkin() error {
	req = buildRequest()
	cookies = buildCookies(viper.GetString("cookie"))
	timeout := viper.GetInt("timeout") //Set the request expiration time
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(req.URL, cookies)
	client := &http.Client{Jar: jar, Timeout: time.Duration(timeout) * time.Second}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	} else {
		return fmt.Errorf("response is nil")
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	body := &resBody{}
	err = json.Unmarshal(response, body)
	Log(fmt.Sprintf("check finish with res  %s \n", body.Message))
	if err != nil {
		return err
	}
	return nil
}

func buildRequest() *http.Request {
	token := "glados.network"
	post := "{\"token\":\"" + token + "\"}"
	reader := strings.NewReader(post)
	request, _ := http.NewRequest(http.MethodPost, "https://glados.rocks/api/user/checkin", reader)
	request.Header.Add("content-type", "application/json; charset=utf-8")
	return request
}

// Parse the string for the cookie
func buildCookies(tmp string) []*http.Cookie {
	if tmp == "" {
		panic("cookie is nil")
	}
	results := strings.Split(tmp, ";")
	cookies := make([]*http.Cookie, len(results))
	for idx, result := range results {
		space := strings.TrimSpace(result)
		kv := strings.Split(space, "=")
		key := kv[0]
		value := kv[1]
		cookie := &http.Cookie{
			Name:  key,
			Value: value,
		}
		cookies[idx] = cookie
	}
	return cookies
}

func configInit() error {
	viper.SetDefault("execTime", DefaultExecTime)
	viper.SetDefault("cookie", "")
	viper.SetDefault("timeout", DefaultTimeout)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return nil
}
