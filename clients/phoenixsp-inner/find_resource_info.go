package phoenixsp_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindResourceInfoRequest struct {
	Caller         string `position:"Query" name:"Caller"`
	InstanceIdList string `position:"Query" name:"InstanceIdList"`
	Aliuid         int64  `position:"Query" name:"Aliuid"`
	Bid            string `position:"Query" name:"Bid"`
	ResourceType   string `position:"Query" name:"ResourceType"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *FindResourceInfoRequest) Invoke(client goaliyun.Client) (*FindResourceInfoResponse, error) {
	resp := &FindResourceInfoResponse{}
	err := client.Request("phoenixsp-inner", "FindResourceInfo", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindResourceInfoResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Count     goaliyun.Integer
	Datas     FindResourceInfoDataList
}

type FindResourceInfoData struct {
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

type FindResourceInfoDataList []FindResourceInfoData

func (list *FindResourceInfoDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindResourceInfoData)
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
