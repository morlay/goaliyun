package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveBatchTaskForCreatingOrderRenewRequest struct {
	UserClientIp     string                                                 `position:"Query" name:"UserClientIp"`
	OrderRenewParams *SaveBatchTaskForCreatingOrderRenewOrderRenewParamList `position:"Query" type:"Repeated" name:"OrderRenewParam"`
	Lang             string                                                 `position:"Query" name:"Lang"`
	RegionId         string                                                 `position:"Query" name:"RegionId"`
}

func (req *SaveBatchTaskForCreatingOrderRenewRequest) Invoke(client goaliyun.Client) (*SaveBatchTaskForCreatingOrderRenewResponse, error) {
	resp := &SaveBatchTaskForCreatingOrderRenewResponse{}
	err := client.Request("domain", "SaveBatchTaskForCreatingOrderRenew", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchTaskForCreatingOrderRenewOrderRenewParam struct {
	DomainName            string `name:"DomainName"`
	CurrentExpirationDate int64  `name:"CurrentExpirationDate"`
	SubscriptionDuration  int64  `name:"SubscriptionDuration"`
}

type SaveBatchTaskForCreatingOrderRenewResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveBatchTaskForCreatingOrderRenewOrderRenewParamList []SaveBatchTaskForCreatingOrderRenewOrderRenewParam

func (list *SaveBatchTaskForCreatingOrderRenewOrderRenewParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SaveBatchTaskForCreatingOrderRenewOrderRenewParam)
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
