package market

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryMarketImagesRequest struct {
	Param    string `position:"Query" name:"Param"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryMarketImagesRequest) Invoke(client goaliyun.Client) (*QueryMarketImagesResponse, error) {
	resp := &QueryMarketImagesResponse{}
	err := client.Request("market", "QueryMarketImages", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMarketImagesResponse struct {
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Result     QueryMarketImagesImageProductList
}

type QueryMarketImagesImageProduct struct {
	ImageProductCode goaliyun.String
	ProductName      goaliyun.String
	CategoryName     goaliyun.String
	SupplierName     goaliyun.String
	BaseSystem       goaliyun.String
	OsKind           goaliyun.String
	OsBit            goaliyun.Integer
	PictureUrl       goaliyun.String
	SmallPicUrl      goaliyun.String
	ShortDescription goaliyun.String
	AgreementUrl     goaliyun.String
	DetailUrl        goaliyun.String
	BuyUrl           goaliyun.String
	StoreUrl         goaliyun.String
	Score            goaliyun.Float
	UserCount        goaliyun.Integer
	SupportIO        bool
	CreatedOn        goaliyun.Integer
	Images           QueryMarketImagesImageList
	SkuCodes         QueryMarketImagesSkuCodeList
	Quota            QueryMarketImagesQuota
	PriceInfo        QueryMarketImagesPriceInfo
}

type QueryMarketImagesImage struct {
	Version            goaliyun.String
	VersionDescription goaliyun.String
	ImageId            goaliyun.String
	ImageSize          goaliyun.Integer
	Region             goaliyun.String
	IsDefault          bool
	SupportIO          bool
	DiskDeviceMappings QueryMarketImagesDiskDeviceMappingList
}

type QueryMarketImagesDiskDeviceMapping struct {
	DiskType        goaliyun.String
	Format          goaliyun.String
	SnapshotId      goaliyun.String
	Size            goaliyun.Integer
	Device          goaliyun.String
	ImportOSSBucket goaliyun.String
	ImportOSSObject goaliyun.String
}

type QueryMarketImagesQuota struct {
	TotalQuota  goaliyun.Integer
	UsingQuota  goaliyun.Integer
	UnusedQuota goaliyun.Integer
}

type QueryMarketImagesPriceInfo struct {
	Rules QueryMarketImagesRuleList
	Order QueryMarketImagesOrder
}

type QueryMarketImagesRule struct {
	RuleId goaliyun.Integer
	Title  goaliyun.String
	Name   goaliyun.String
}

type QueryMarketImagesOrder struct {
	OriginalPrice goaliyun.Float
	DiscountPrice goaliyun.Float
	TradePrice    goaliyun.Float
	Currency      goaliyun.String
	Period        goaliyun.Integer
	PriceUnit     goaliyun.String
	RuleIdSet     QueryMarketImagesRuleIdSetList
}

type QueryMarketImagesImageProductList []QueryMarketImagesImageProduct

func (list *QueryMarketImagesImageProductList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketImagesImageProduct)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type QueryMarketImagesImageList []QueryMarketImagesImage

func (list *QueryMarketImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketImagesImage)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type QueryMarketImagesSkuCodeList []goaliyun.String

func (list *QueryMarketImagesSkuCodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type QueryMarketImagesDiskDeviceMappingList []QueryMarketImagesDiskDeviceMapping

func (list *QueryMarketImagesDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketImagesDiskDeviceMapping)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type QueryMarketImagesRuleList []QueryMarketImagesRule

func (list *QueryMarketImagesRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketImagesRule)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type QueryMarketImagesRuleIdSetList []goaliyun.String

func (list *QueryMarketImagesRuleIdSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
