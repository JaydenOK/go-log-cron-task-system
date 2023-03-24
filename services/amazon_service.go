package services

import (
	"app/contexts"
	"app/libs/mysqllib"
	"app/models"
	"app/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
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
	fmt.Println(lists)
	return lists, nil

	taskCh := make(chan int, concurrencyInt)
	var total int
	for i, account := range lists {
		taskCh <- i
		total++
		go service.doRefreshToken(taskCh, account)
	}
	data := make(map[string]interface{})
	data["total"] = total
	return data, nil
}

//执行任务
func (service *AmazonService) doRefreshToken(taskCh chan int, model models.AmazonAdsAccountModel) {

	<-taskCh
}

func getAccountList(db *gorm.DB, page, pageSize int) []models.AmazonAdsAccountModel {
	var amazonAdsAccountModel []models.AmazonAdsAccountModel //用于查找多个
	db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id asc").Find(&amazonAdsAccountModel)
	return amazonAdsAccountModel
}

func (service *AmazonService) PullOrder(ctx contexts.MyContext) interface{} {
	return nil
}

func (service *AmazonService) GetReportDocument(ctx contexts.MyContext) interface{} {
	return nil
}
