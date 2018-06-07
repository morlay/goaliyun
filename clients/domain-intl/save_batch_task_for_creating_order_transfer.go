package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveBatchTaskForCreatingOrderTransferRequest struct {
	OrderTransferParams *SaveBatchTaskForCreatingOrderTransferOrderTransferParamList `position:"Query" type:"Repeated" name:"OrderTransferParam"`
	UserClientIp        string                                                       `position:"Query" name:"UserClientIp"`
	Lang                string                                                       `position:"Query" name:"Lang"`
	RegionId            string                                                       `position:"Query" name:"RegionId"`
}

func (req *SaveBatchTaskForCreatingOrderTransferRequest) Invoke(client goaliyun.Client) (*SaveBatchTaskForCreatingOrderTransferResponse, error) {
	resp := &SaveBatchTaskForCreatingOrderTransferResponse{}
	err := client.Request("domain-intl", "SaveBatchTaskForCreatingOrderTransfer", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchTaskForCreatingOrderTransferOrderTransferParam struct {
	DomainName            string `name:"DomainName"`
	AuthorizationCode     string `name:"AuthorizationCode"`
	RegistrantProfileId   int64  `name:"RegistrantProfileId"`
	PermitPremiumTransfer string `name:"PermitPremiumTransfer"`
}

type SaveBatchTaskForCreatingOrderTransferResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveBatchTaskForCreatingOrderTransferOrderTransferParamList []SaveBatchTaskForCreatingOrderTransferOrderTransferParam

func (list *SaveBatchTaskForCreatingOrderTransferOrderTransferParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SaveBatchTaskForCreatingOrderTransferOrderTransferParam)
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
