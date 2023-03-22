package utils

import (
	"encoding/json"
	"net/http"
)

//获取参数方式
//方式一：r.Header["Accept-Encoding"]，得到的是一个字符串切片
//方式二：r.Header.Get("Accept-Encoding")，得到的是字符串形式的值，多个值使用逗号分隔
func ParamsGet(r *http.Request) interface{} {
	query := r.URL.Query()
	return query
}

func ParamsPost(r *http.Request) interface{} {
	r.ParseForm()
	return r.PostForm
}

func ParamsJson(r *http.Request) interface{} {
	// 根据请求body创建一个json解析器实例
	decoder := json.NewDecoder(r.Body)
	// params用于存放参数key=value数据，或解析到结构体
	var params map[string]string
	// 解析参数 存入map
	decoder.Decode(&params)
	return params
}

func RequestHeader(r *http.Request) interface{} {
	return r.Header
}

func HttpMethod(r *http.Request) interface{} {
	return r.Method
}
