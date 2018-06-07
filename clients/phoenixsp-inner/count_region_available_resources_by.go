package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type CountRegionAvailableResourcesByRequest struct {
	MarketType         string `position:"Query" name:"MarketType"`
	Filter             string `position:"Query" name:"Filter"`
	Caller             string `position:"Query" name:"Caller"`
	ResCreateTimeEnd   int64  `position:"Query" name:"ResCreateTimeEnd"`
	ResourceSource     string `position:"Query" name:"ResourceSource"`
	AliUID             int64  `position:"Query" name:"AliUID"`
	ChargeType         string `position:"Query" name:"ChargeType"`
	CommodityCode      string `position:"Query" name:"CommodityCode"`
	Bid                string `position:"Query" name:"Bid"`
	Region             string `position:"Query" name:"Region"`
	ResourceType       string `position:"Query" name:"ResourceType"`
	ResCreateTimeBegin int64  `position:"Query" name:"ResCreateTimeBegin"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *CountRegionAvailableResourcesByRequest) Invoke(client goaliyun.Client) (*CountRegionAvailableResourcesByResponse, error) {
	resp := &CountRegionAvailableResourcesByResponse{}
	err := client.Request("phoenixsp-inner", "CountRegionAvailableResourcesBy", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CountRegionAvailableResourcesByResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
	Data      goaliyun.Integer
}
