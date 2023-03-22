package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	StatusSuccess = 1
	StatusFail    = 9999
)

type ResponseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//成功响应
func SuccessResponse(w http.ResponseWriter, data interface{}) {
	JsonResponse(w, http.StatusOK, StatusSuccess, "success", data)
}

//失败响应
func FailResponse(w http.ResponseWriter, data interface{}) {
	JsonResponse(w, http.StatusOK, StatusFail, "fail", data)
}

//信息响应
func MessageResponse(w http.ResponseWriter, message string) {
	JsonResponse(w, http.StatusOK, StatusSuccess, "success", message)
}

//重定向
func Redirect(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
	w.WriteHeader(302)
}

//http-json响应
func JsonResponse(w http.ResponseWriter, httpCode, status int, message string, data interface{}) {
	responseData := ResponseData{
		Status:  status,
		Message: message,
		Data:    data,
	}
	bytes, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}
	//响应返回客户端json数据
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Token", "aaaa")
	//WriteHeader放着Set后
	w.WriteHeader(httpCode)
	if i, err := w.Write(bytes); err != nil {
		fmt.Println("JsonResponse Fail:", i)
	}
}
