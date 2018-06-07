package phoenixsp_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindUserAvailableResourcesByRequest struct {
	BeginEndTime       int64  `position:"Query" name:"BeginEndTime"`
	EndResCreateTime   int64  `position:"Query" name:"EndResCreateTime"`
	BillingTag         string `position:"Query" name:"BillingTag"`
	CommodityCode      string `position:"Query" name:"CommodityCode"`
	AvailableDays      int64  `position:"Query" name:"AvailableDays"`
	ResourceType       string `position:"Query" name:"ResourceType"`
	BeginResCreateTime int64  `position:"Query" name:"BeginResCreateTime"`
	CurrPage           int64  `position:"Query" name:"CurrPage"`
	Caller             string `position:"Query" name:"Caller"`
	AutoRenewal        string `position:"Query" name:"AutoRenewal"`
	EndEndTime         int64  `position:"Query" name:"EndEndTime"`
	InstanceIdList     string `position:"Query" name:"InstanceIdList"`
	RenewalStatus      string `position:"Query" name:"RenewalStatus"`
	HasExpiredRes      string `position:"Query" name:"HasExpiredRes"`
	PageSize           int64  `position:"Query" name:"PageSize"`
	Aliuid             int64  `position:"Query" name:"Aliuid"`
	IsPrepaid          string `position:"Query" name:"IsPrepaid"`
	Bid                string `position:"Query" name:"Bid"`
	Region             string `position:"Query" name:"Region"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *FindUserAvailableResourcesByRequest) Invoke(client goaliyun.Client) (*FindUserAvailableResourcesByResponse, error) {
	resp := &FindUserAvailableResourcesByResponse{}
	err := client.Request("phoenixsp-inner", "FindUserAvailableResourcesBy", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindUserAvailableResourcesByResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Count     goaliyun.Integer
	Datas     FindUserAvailableResourcesByDataList
}

type FindUserAvailableResourcesByData struct {
	Aliuid              goaliyun.Integer
	BuyerId             goaliyun.Integer
	InstanceId          goaliyun.String
	Region              goaliyun.String
	ResourceType        goaliyun.String
	ChargeType          goaliyun.String
	EndTime             goaliyun.Integer
	ReleaseTime         goaliyun.Integer
	ExtraEndTime        goaliyun.Integer
	ResCreateTime       goaliyun.Integer
	BillingTag          goaliyun.String
	CommodityCode       goaliyun.String
	ResourceStatus      goaliyun.String
	ResourceSubStatus   goaliyun.String
	ExpectedReleaseTime goaliyun.Integer
	Bid                 goaliyun.String
	AutoRenewal         bool
	RenewalStatus       goaliyun.String
	RenewalDuration     goaliyun.Integer
	RenewalCycUnit      goaliyun.Integer
	SaleCycle           goaliyun.String
	MarketType          goaliyun.String
}

type FindUserAvailableResourcesByDataList []FindUserAvailableResourcesByData

func (list *FindUserAvailableResourcesByDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindUserAvailableResourcesByData)
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
