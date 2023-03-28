package services

import (
	"app/contexts"
	"app/libs/httplib"
	"app/libs/loglib"
	"app/libs/mysqllib"
	"app/models"
	"app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type AmazonService struct {
}

type taskLock struct {
	mutex sync.Mutex
}

//定时任务触发，并发刷新亚马逊token
func (service *AmazonService) RefreshToken(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	query := utils.ParamsGet(r)
	//并发数
	concurrency := query.Get("concurrency")
	id := query.Get("id")
	var concurrencyInt, idInt int
	var e error
	concurrencyInt, e = utils.StringToInt(concurrency)
	if e != nil {
		return nil, errors.New("参数concurrency错误")
	}
	if id != "" {
		idInt, e = utils.StringToInt(concurrency)
		if e != nil {
			return nil, errors.New("参数id错误")
		}
	}
	pageSize := 10
	lists := getAccountList(1, pageSize, idInt)
	if lists == nil {
		return map[string]string{"message": "nodata"}, nil
	}
	//获取开发者信息
	developerList := getDeveloperList()
	fmt.Println("developerList", developerList)
	//@todo 使用chan有限通道，阻塞投递，达到限制并发目的，控制任务并发数
	loglib.Info("concurrency:" + utils.IntToString(concurrencyInt))
	//@todo 主协程阻塞
	var wg sync.WaitGroup
	//@todo 任务并发阻塞（匿名空struct）
	var taskCh chan struct{}
	taskCh = make(chan struct{}, concurrencyInt)
	defer close(taskCh)
	var total int
	for _, account := range lists {
		wg.Add(1)
		//阻塞投递
		taskCh <- struct{}{}
		var developer models.AmazonDeveloperModel
		for _, value := range developerList {
			if value.AppId == account.AdsAppId {
				developer = value
			}
		}
		total++
		go service.doRefreshToken(&wg, taskCh, account, developer)
	}
	wg.Wait()
	data := make(map[string]interface{})
	data["total"] = total
	return data, nil
}

//执行任务：
func (service *AmazonService) doRefreshToken(wg *sync.WaitGroup, taskCh chan struct{},
	account models.AmazonAdsAccountModel, developer models.AmazonDeveloperModel) {
	defer func() {
		<-taskCh
	}()
	defer wg.Done()
	db := mysqllib.GetMysqlDb()
	requestUrl := "https://api.amazon.com/auth/o2/token"
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", account.AdsRefreshToken)
	data.Set("client_id", developer.ClientId)
	data.Set("client_secret", developer.ClientSecret)
	resp := httplib.PostForm(requestUrl, data)
	var tokenData map[string]interface{}
	err := json.Unmarshal([]byte(resp), &tokenData)
	if err != nil {
		// 根据条件更新
		db.Model(&models.AmazonAdsAccountModel{}).Where("id=?", account.Id).Update("ads_refresh_msg", resp)
		//db.Model(&amazonAdsAccountModel).Where("id=?", account.Id).Update("ads_refresh_msg", resp)  //&struct取实例化后，使用实例化的主键作为条件更新
		loglib.Info("fail:" + resp)
		return
	}
	_, ok := tokenData["access_token"]
	if !ok {
		value, ok2 := tokenData["error_description"]
		if ok2 {
			db.Model(&models.AmazonAdsAccountModel{}).Where("id=?", account.Id).Update("ads_refresh_msg", value)
		} else {
			db.Model(&models.AmazonAdsAccountModel{}).Where("id=?", account.Id).Update("ads_refresh_msg", resp)
		}
		return
	}
	expiresIn := tokenData["expires_in"]
	authorizeEndTime := utils.GetCurrentTimestamp() + int64(utils.InterfaceToInt(expiresIn))
	//Updates更新多列
	set := map[string]interface{}{
		"ads_access_token":       tokenData["access_token"],
		"ads_authorize_end_time": utils.FormatTimeToDate(authorizeEndTime),
		"ads_refresh_token":      tokenData["refresh_token"],
		"ads_expires_in":         tokenData["expires_in"],
		"ads_refresh_time":       utils.GetCurrentDateTime(),
		"ads_refresh_msg":        "",
	}
	result := db.Model(&models.AmazonAdsAccountModel{}).Where("id=?", account.Id).Updates(set)
	if result.RowsAffected <= 0 {
		loglib.Info("update fail:" + strconv.Itoa(int(result.RowsAffected)))
		return
	}
	loglib.Info("Done| id:" + utils.Int32ToString(account.Id) + ", account_name:" + account.AccountName + ", resp:" + resp)
	return
}

func getAccountList(page, pageSize, idInt int) []models.AmazonAdsAccountModel {
	mysqlDb := mysqllib.GetMysqlDb()
	db := mysqlDb.Table((&models.AmazonAdsAccountModel{}).TableName())
	if idInt != 0 {
		db.Where("id=?", idInt)
	}
	var amazonAdsAccountModels []models.AmazonAdsAccountModel //多个切片
	db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id asc").Find(&amazonAdsAccountModels)
	return amazonAdsAccountModels
}

func getDeveloperList() []models.AmazonDeveloperModel {
	dbSystem := mysqllib.GetMysqlDbSystem()
	var amazonDeveloperModel []models.AmazonDeveloperModel //多个切片
	dbSystem.Find(&amazonDeveloperModel)
	return amazonDeveloperModel
}

func (service *AmazonService) PullOrder(ctx contexts.MyContext) interface{} {
	return nil
}

func (service *AmazonService) GetReportDocument(ctx contexts.MyContext) interface{} {
	return nil
}
