package services

import (
	"app/contexts"
	"app/utils"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type AmazonService struct {
}

//定时任务触发，并发刷新亚马逊token
func (service *AmazonService) RefreshToken(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	utils.ParamsGet(r)

	var users map[string]User
	users = make(map[string]User) //在使用时没有进行初始化map，导致使用时失败，或者直接声明时，使用
	userListJson, _ := redisClient.Get(UserListKey).Result()
	_ = json.Unmarshal([]byte(userListJson), &users)
	// 获取 map 中某个 key 是否存在的语法。如果 ok 是 true，表示 key 存在，key 对应的值就是 value ，反之表示 key 不存在。
	_, ok := users[username]
	if !ok {
		users[username] = User{
			Username: username,
			Password: password,
		}
		newUserListJson, _ := json.Marshal(users)
		redisClient.Set(UserListKey, newUserListJson, 86400*time.Second)
		return users, nil
	} else {
		return nil, errors.New("用户已注册")
	}
}

func (service *AmazonService) PullOrder(ctx contexts.MyContext) interface{} {
	return nil
}

func (service *AmazonService) GetReportDocument(ctx contexts.MyContext) interface{} {
	return nil
}
