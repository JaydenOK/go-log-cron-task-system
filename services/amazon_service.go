package services

import (
	"app/contexts"
	"app/libs/loglib"
	"app/libs/mysqllib"
	"app/models"
	"app/utils"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"sync"
	"time"
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
		total++
		go service.doRefreshToken(&wg, taskCh, account)
	}
	wg.Wait()
	data := make(map[string]interface{})
	data["total"] = total
	return data, nil
}

//执行任务
func (service *AmazonService) doRefreshToken(wg *sync.WaitGroup, taskCh chan struct{}, account models.AmazonAdsAccountModel) {
	defer func() {
		<-taskCh
	}()
	defer wg.Done()
	time.Sleep(5 * time.Second)
	loglib.Info("task done:" + utils.Int32ToString(account.Id) + ", account_name:" + account.AccountName)
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
