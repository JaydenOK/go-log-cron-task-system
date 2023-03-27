package httplib

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//结构体参数
//type option struct {
//	Url         string
//	Data        interface{}
//	Timeout     int
//	ContentType string
//	Cookie      string
//	Header      map[string]string
//}

var defaultTimeOut = time.Second * 10

//发送PostForm参数，data map[string][]string
//data := url.Values{}
//data.Set("grant_type", "refresh_token")
//resp := httplib.PostForm(requestUrl, data)
func PostForm(requestUrl string, data url.Values) string {
	client := &http.Client{
		Timeout: defaultTimeOut,
	}
	resp, err := client.Post(requestUrl, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Printf("post failed, err:%v", err)
		return ""
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v", err)
		return ""
	}
	return string(b)
}

//
//func PostForm(url string, data url.Values) (string, error) {
//	resp, err := http.PostForm(url, data)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//	content, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//	return string(content), nil
//}

//发送json参数
func PostJson(requestUrl string, data string) string {
	client := &http.Client{
		Timeout: defaultTimeOut,
	}
	resp, err := client.Post(requestUrl, "application/json", strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v", err)
		return ""
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v", err)
		return ""
	}
	return string(b)
}

//第三方rest http方法
func demo() {
	client := resty.New()
	// POST Map, default is JSON content type. No need to set one
	resp, err := client.R().
		SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
		SetResult(&struct{}{}). // or SetResult(AuthSuccess{}).
		SetError(&struct{}{}). // or SetError(AuthError{}).
		Post("https://myapp.com/login")
	fmt.Println(resp, err)
}
