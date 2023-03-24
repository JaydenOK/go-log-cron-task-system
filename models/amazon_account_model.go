package models

import "app/utils"

type AmazonAdsAccountModel struct {
	Id                  int32           `json:"id"`
	AccountId           int32           `json:"account_id"`
	AccountName         string           `json:"account_name"`
	OrderId             string          `json:"order_id"`
	SellingPartnerId    string          `json:"selling_partner_id"`
	ProfileId           string          `json:"profile_id"`
	SiteId              string          `json:"site_id"`
	AdsAppId            string          `json:"ads_app_id"`
	AdsAccessToken      string          `json:"ads_access_token"`
	AdsAuthorizeEndTime string          `json:"ads_authorize_end_time"`
	AdsExpiresIn        string          `json:"ads_expires_in"`
	AdsRefreshToken     string          `json:"ads_refresh_token"`
	CreateTime          utils.LocalTime `json:"create_time"` //utils.LocalTime： 实现MarshalJSON接口，格式化数据
}

func (amazonAdsAccountModel *AmazonAdsAccountModel) TableName() string {
	return "yibai_amazon_ads_account"
}
