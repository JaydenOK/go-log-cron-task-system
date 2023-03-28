package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"
const DateFormat = "2006-01-02"

type LocalTime time.Time

// 实现MarshalJSON接口，格式化数据，解决 c.JSON 时解析值的问题
func (t LocalTime) MarshalJSON() ([]byte, error) {
	if &t == nil {
		return []byte("null"), nil
	}
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(TimeFormat))
	return []byte(stamp), nil
}

// 在 c.ShouldBindJSON 时，会调用 field.UnmarshalJSON 方法
func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}
	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

// 写入 mysql 时调用
func (t LocalTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

// 检出 mysql 时调用
func (t *LocalTime) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

// 获取当前时间戳 1664531446
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// 获取毫秒时间戳 1664531446277
func GetCurrentMilliTimestamp() int64 {
	return time.Now().UnixMilli()
}

// 获取毫秒时间戳 1664531446277667
func GetCurrentMicroTimestamp() int64 {
	return time.Now().UnixMicro()
}

// 获取当前日期时间 2022-09-30 17:50:46
func GetCurrentDateTime() string {
	return time.Now().Format(TimeFormat)
}

// 获取当前日期 2022-09-30
func GetCurrentDate() string {
	return time.Now().Format(DateFormat)
}

//时间戳转日期(11位)
func FormatTimeToDate(t int64) string {
	return time.Unix(t, 0).Format(TimeFormat)
}

//日期字符串转时间戳
func ParseDateToTime(s string) int64 {
	if t, e := time.Parse(TimeFormat, s); e != nil {
		return 0
	} else {
		return t.Unix()
	}
}
