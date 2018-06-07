package phoenixsp_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindBillingInfoByUserInPeriodRequest struct {
	CurrPage        int64  `position:"Query" name:"CurrPage"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	EndTime         int64  `position:"Query" name:"EndTime"`
	AliUID          int64  `position:"Query" name:"AliUID"`
	CommodityCode   string `position:"Query" name:"CommodityCode"`
	StartTime       int64  `position:"Query" name:"StartTime"`
	Bid             string `position:"Query" name:"Bid"`
	TradeInstanceID string `position:"Query" name:"TradeInstanceID"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *FindBillingInfoByUserInPeriodRequest) Invoke(client goaliyun.Client) (*FindBillingInfoByUserInPeriodResponse, error) {
	resp := &FindBillingInfoByUserInPeriodResponse{}
	err := client.Request("phoenixsp-inner", "FindBillingInfoByUserInPeriod", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindBillingInfoByUserInPeriodResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Count     goaliyun.Integer
	Datas     FindBillingInfoByUserInPeriodDataList
}

type FindBillingInfoByUserInPeriodData struct {
	AliUID           goaliyun.Integer
	ResourceType     goaliyun.String
	ResourceStatus   goaliyun.String
	InstanceId       goaliyun.String
	BillingTag       goaliyun.String
	ChargeType       goaliyun.String
	ResourceSource   goaliyun.String
	CommodityCode    goaliyun.String
	PropertyDetail   goaliyun.String
	Operation        goaliyun.String
	CommandStartTime goaliyun.Integer
	ActualStartTime  goaliyun.Integer
	ExpectStartTime  goaliyun.Integer
	CommandEndTime   goaliyun.Integer
	ActualEndTime    goaliyun.Integer
	ExpectEndTime    goaliyun.Integer
}

type FindBillingInfoByUserInPeriodDataList []FindBillingInfoByUserInPeriodData

func (list *FindBillingInfoByUserInPeriodDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindBillingInfoByUserInPeriodData)
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
