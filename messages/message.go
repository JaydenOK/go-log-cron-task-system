package messages

//在 Golang 的结构体定义中添加 omitempty 关键字，来表示这条信息如果没有提供，在序列化成 json 字符串的时候就不要包含其默认值

type SearchMessage struct {
	Id         string `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time,omitempty"`
}
