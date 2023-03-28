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
func InterfaceToString(f interface{}) string {
	// interface 转 string
	var s string
	if f == nil {
		return s
	}
	switch f.(type) {
	case float64:
		ft := f.(float64)
		s = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := f.(float32)
		s = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := f.(int)
		s = strconv.Itoa(it)
	case uint:
		it := f.(uint)
		s = strconv.Itoa(int(it))
	case int8:
		it := f.(int8)
		s = strconv.Itoa(int(it))
	case uint8:
		it := f.(uint8)
		s = strconv.Itoa(int(it))
	case int16:
		it := f.(int16)
		s = strconv.Itoa(int(it))
	case uint16:
		it := f.(uint16)
		s = strconv.Itoa(int(it))
	case int32:
		it := f.(int32)
		s = strconv.Itoa(int(it))
	case uint32:
		it := f.(uint32)
		s = strconv.Itoa(int(it))
	case int64:
		it := f.(int64)
		s = strconv.FormatInt(it, 10)
	case uint64:
		it := f.(uint64)
		s = strconv.FormatUint(it, 10)
	case string:
		s = f.(string)
	case []byte:
		s = string(f.([]byte))
	default:
		newValue, _ := json.Marshal(f)
		s = string(newValue)
	}
	return s
}

//interface转Int
func InterfaceToInt(f interface{}) int {
	var i int
	switch f.(type) {
	case uint:
		i = int(f.(uint))
		break
	case int8:
		i = int(f.(int8))
		break
	case uint8:
		i = int(f.(uint8))
		break
	case int16:
		i = int(f.(int16))
		break
	case uint16:
		i = int(f.(uint16))
		break
	case int32:
		i = int(f.(int32))
		break
	case uint32:
		i = int(f.(uint32))
		break
	case int64:
		i = int(f.(int64))
		break
	case uint64:
		i = int(f.(uint64))
		break
	case float32:
		i = int(f.(float32))
		break
	case float64:
		i = int(f.(float64))
		break
	case string:
		i, _ = strconv.Atoi(f.(string))
		if i == 0 && len(f.(string)) > 0 {
			f, _ := strconv.ParseFloat(f.(string), 64)
			i = int(f)
		}
		break
	case nil:
		i = 0
		break
	case json.Number:
		t3, _ := f.(json.Number).Int64()
		i = int(t3)
		break
	default:
		i = f.(int)
		break
	}
	return i
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
