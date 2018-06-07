package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type CountUserAvailableResourcesByRequest struct {
	BeginEndTime       int64  `position:"Query" name:"BeginEndTime"`
	EndResCreateTime   int64  `position:"Query" name:"EndResCreateTime"`
	BillingTag         string `position:"Query" name:"BillingTag"`
	CommodityCode      string `position:"Query" name:"CommodityCode"`
	AvailableDays      int64  `position:"Query" name:"AvailableDays"`
	ResourceType       string `position:"Query" name:"ResourceType"`
	BeginResCreateTime int64  `position:"Query" name:"BeginResCreateTime"`
	Caller             string `position:"Query" name:"Caller"`
	AutoRenewal        string `position:"Query" name:"AutoRenewal"`
	EndEndTime         int64  `position:"Query" name:"EndEndTime"`
	InstanceIdList     string `position:"Query" name:"InstanceIdList"`
	HasExpiredRes      string `position:"Query" name:"HasExpiredRes"`
	Aliuid             int64  `position:"Query" name:"Aliuid"`
	IsPrepaid          string `position:"Query" name:"IsPrepaid"`
	Bid                string `position:"Query" name:"Bid"`
	Region             string `position:"Query" name:"Region"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *CountUserAvailableResourcesByRequest) Invoke(client goaliyun.Client) (*CountUserAvailableResourcesByResponse, error) {
	resp := &CountUserAvailableResourcesByResponse{}
	err := client.Request("phoenixsp-inner", "CountUserAvailableResourcesBy", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CountUserAvailableResourcesByResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
	Data      goaliyun.Integer
}
