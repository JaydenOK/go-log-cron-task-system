package models

type AmazonAdsSiteModel struct {
	SiteId     uint32 `json:"site_id"`
	SiteCode   string `json:"site_code"`
	SiteUrl    string `json:"site_url"`
	SiteRegion string `json:"site_region"`
}

func (*AmazonAdsSiteModel) TableName() string {
	return "yibai_amazon_ads_site"
}
