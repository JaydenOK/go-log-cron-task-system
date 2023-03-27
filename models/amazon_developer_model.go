package models

import "app/utils"

type AmazonDeveloperModel struct {
	AppId         int32           `json:"app_id"`
	ApplicationId string           `json:"application_id"`
	ClientId      string          `json:"client_id"`
	ClientSecret  string          `json:"client_secret"`
	CreateTime    utils.LocalTime `json:"create_time"`
}

func (amazonDeveloperModel *AmazonDeveloperModel) TableName() string {
	return "yibai_amazon_ads_developer"
}
