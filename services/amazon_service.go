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
	"gorm.io/gorm"
	"net/http"
	"net/url"
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
	mysqlDb := mysqllib.GetMysqlDb()
	var amazonAdsAccountModel models.AmazonAdsAccountModel
	db := mysqlDb.Table(amazonAdsAccountModel.TableName())
	if idInt != 0 {
		db = db.Where("id=?", idInt)
	}
	pageSize := 10
	lists := getAccountList(db, 1, pageSize)
	if lists == nil {
		return map[string]string{"message": "nodata"}, nil
	}
	//获取开发者信息
	dbSystem := mysqllib.GetMysqlDbSystem()
	developerList := getDeveloperList(dbSystem)
	fmt.Println(developerList)
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
		go service.doRefreshToken(db, &wg, taskCh, account, developer)
	}
	wg.Wait()
	data := make(map[string]interface{})
	data["total"] = total
	return data, nil
}

//执行任务：
func (service *AmazonService) doRefreshToken(db *gorm.DB, wg *sync.WaitGroup, taskCh chan struct{},
	account models.AmazonAdsAccountModel, developer models.AmazonDeveloperModel) string {
	defer func() {
		<-taskCh
	}()
	defer wg.Done()
	//requestUrl := "http://192.168.92.208/shop/Auth/code"
	requestUrl := "https://api.amazon.com/auth/o2/token"
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", account.AdsRefreshToken)
	data.Set("client_id", developer.ClientId)
	data.Set("client_secret", developer.ClientSecret)
	resp := httplib.PostForm(requestUrl, data)
	//{"access_token":"Atza|IwEBIBRHABUnTDIu03Kqdrpp9J3XCwFeNqMTSSdCEsEN10ZMYlyoZT3xzPnJ1pud6eQtppsFrK-GknMZt0PL3BS7ikdSQwcoLCc8TfWXQo_UWfmXevkwer7GJhIjkBMOuWnq8WJ5B-cvU0Tc0u_8bfDd05ny6rKRfryJAbuxVYF5Te_3kXNqjwLopK_igZNy59EJMOTjG22kl01OoY_jWp92xGbjjvCsY7-DvWRJh_tnuLLFHFERET0tIBapiPamALx4s4VKqUHrYtfH8ixuqPRXFPqLZtqvV2wjQJ2BKtTuy5Ssy6He5mKomeVS8ktQa9n6BjPaDaU2XyS55NUkRp5WzVmuJeArGrKuWK7yW-u1jarqzM1Z6xkdlkyFJM9l8zWcnYNHleX1IcgenRqqAyLt4yzNnGAySR9KThu0hWODKn524Vuo3V6OEvJWxPNt_d2jBe_AcH0L1rCTDhJxue7Xddd7HPVkMOJME7QdULIlkvhluw","refresh_token":"Atzr|IwEBIEvJPwllpNsOq60juN5SQl6zpi7PjZ8WtANR-LekhiuCDlVLj-MI-DrszKxr5xGsrmkwp2PIhoA0wWLG-eQVkXKbmRu-WkECRqIH4f2m0Ve8p3zmx9AZoid01V69NCBxVJA9CxbB-l0uQh2m4SJ3s7f4cKX2tk-pIFN2X2B4b5ZWna5MMt3RNsmMiEko7XZITftAXh0R9rBdbocVcZb_pvXxHvORYXpkiFLKsxcew8BIBCglU1BiGTV-7UdAfD1NOpNv4FSZNOLdaLNAnLwQuc4OuBJTtyAtfcus8GQ0imxQdo6UzcRo9D3TOoxsmjRc0rIfa31x83buaOmZpqvLgwpqkwQ1MMyA1dXOhC7A0lRyMZ2A1Hpki5xVqT9uUR72K2T7WbLehskYfGWGRocML0UDs1l5Mp-PeNU6lcHe_Ok4kkSz0trqEsXJ2KBu5ZnrhLuv2c3a6ERNsAQBltJQnJPpLIMHM02pnjNrF7VKcn--LQ","token_type":"bearer","expires_in":3600}
	var tokenData map[string]interface{}
	err := json.Unmarshal([]byte(resp), &tokenData)
	if err != nil {
		db.Model(models.AmazonAdsAccountModel{}).Where("id=?", account.Id).Update("ads_refresh_msg", resp)
		return "fail:" + resp
	}
	accessToken := tokenData["access_token"]
	refreshToken := tokenData["refresh_token"]
	//tokenType := tokenData["token_type"]
	expiresIn := tokenData["expires_in"]
	fmt.Println()
	//authorizeEndTime := utils.GetCurrentTimestamp() + expiresIn.(int64)
	//Updates更新多列
	//set := map[string]interface{}{"ads_access_token": accessToken, "ads_authorize_end_time": authorizeEndTime, "ads_refresh_token": refreshToken, "ads_expires_in": expiresIn,}
	set := map[string]interface{}{"ads_access_token": accessToken, "ads_refresh_token": refreshToken, "ads_expires_in": expiresIn,}
	db.Model(models.AmazonAdsAccountModel{}).Where("id=?", account.Id).Updates(set)
	//db.Model(models.AmazonAdsAccountModel{}).Where("id=?", account.Id).Updates(models.AmazonAdsAccountModel{AdsRefreshMsg:resp, AdsAccessToken: ""})
	loglib.Info("Done| id:" + utils.Int32ToString(account.Id) + ", account_name:" + account.AccountName + ", resp:" + resp)
	return resp
}

func getAccountList(db *gorm.DB, page, pageSize int) []models.AmazonAdsAccountModel {
	var amazonAdsAccountModel []models.AmazonAdsAccountModel //多个切片
	db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id asc").Find(&amazonAdsAccountModel)
	return amazonAdsAccountModel
}

func getDeveloperList(db *gorm.DB) []models.AmazonDeveloperModel {
	var amazonDeveloperModel []models.AmazonDeveloperModel //多个切片
	db.Find(&amazonDeveloperModel)
	return amazonDeveloperModel
}

func (service *AmazonService) PullOrder(ctx contexts.MyContext) interface{} {
	return nil
}

func (service *AmazonService) GetReportDocument(ctx contexts.MyContext) interface{} {
	return nil
}
