package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveBatchTaskForCreatingOrderRedeemRequest struct {
	OrderRedeemParams *SaveBatchTaskForCreatingOrderRedeemOrderRedeemParamList `position:"Query" type:"Repeated" name:"OrderRedeemParam"`
	UserClientIp      string                                                   `position:"Query" name:"UserClientIp"`
	Lang              string                                                   `position:"Query" name:"Lang"`
	RegionId          string                                                   `position:"Query" name:"RegionId"`
}

func (req *SaveBatchTaskForCreatingOrderRedeemRequest) Invoke(client goaliyun.Client) (*SaveBatchTaskForCreatingOrderRedeemResponse, error) {
	resp := &SaveBatchTaskForCreatingOrderRedeemResponse{}
	err := client.Request("domain", "SaveBatchTaskForCreatingOrderRedeem", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchTaskForCreatingOrderRedeemOrderRedeemParam struct {
	DomainName            string `name:"DomainName"`
	CurrentExpirationDate int64  `name:"CurrentExpirationDate"`
}

type SaveBatchTaskForCreatingOrderRedeemResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveBatchTaskForCreatingOrderRedeemOrderRedeemParamList []SaveBatchTaskForCreatingOrderRedeemOrderRedeemParam

func (list *SaveBatchTaskForCreatingOrderRedeemOrderRedeemParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SaveBatchTaskForCreatingOrderRedeemOrderRedeemParam)
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
