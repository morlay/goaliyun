package emr

import (
	"github.com/morlay/goaliyun"
)

type QueryPriceForRenewEcsRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	EcsPeriod       int64  `position:"Query" name:"EcsPeriod"`
	EmrPeriod       int64  `position:"Query" name:"EmrPeriod"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	EcsId           string `position:"Query" name:"EcsId"`
	EcsIdList       string `position:"Query" name:"EcsIdList"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryPriceForRenewEcsRequest) Invoke(client goaliyun.Client) (*QueryPriceForRenewEcsResponse, error) {
	resp := &QueryPriceForRenewEcsResponse{}
	err := client.Request("emr", "QueryPriceForRenewEcs", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryPriceForRenewEcsResponse struct {
	RequestId  goaliyun.String
	EcsId      goaliyun.String
	EmrPrice   goaliyun.String
	EcsPrice   goaliyun.String
	EmrPriceDO QueryPriceForRenewEcsEmrPriceDO
	EcsPriceDO QueryPriceForRenewEcsEcsPriceDO
}

type QueryPriceForRenewEcsEmrPriceDO struct {
	OriginalPrice goaliyun.String
	DiscountPrice goaliyun.String
	TradePrice    goaliyun.String
	TaxPrice      goaliyun.String
	Currency      goaliyun.String
}

type QueryPriceForRenewEcsEcsPriceDO struct {
	OriginalPrice goaliyun.String
	DiscountPrice goaliyun.String
	TradePrice    goaliyun.String
	TaxPrice      goaliyun.String
	Currency      goaliyun.String
}
