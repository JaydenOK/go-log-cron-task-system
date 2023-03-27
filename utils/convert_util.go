package utils

import (
	"encoding/json"
	"strconv"
)

// 字符串string转整型
func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// 整型int转字符串
func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int32ToString(i int32) string {
	return strconv.Itoa(int(i))
}

// 字符串转Byte切片
func StringToByte(str string) []byte {
	return []byte(str)
}

// Byte转字符串
func ByteToString(b []byte) string {
	return string(b)
}

// interface转string
func InterfaceToString(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// 字符串转interface
func StringToInterface(s string) interface{} {
	//return s	//直接返回
	var x interface{} = s
	return x
}

//map集合转字节
func MapToByte(m map[string]interface{}) ([]byte, error) {
	bytes, err := json.Marshal(m)
	return bytes, err
}

//字节转map集合
func ByteToMap(b []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	return m, err
}
