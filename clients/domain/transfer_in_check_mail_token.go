package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type TransferInCheckMailTokenRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Token        string `position:"Query" name:"Token"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *TransferInCheckMailTokenRequest) Invoke(client goaliyun.Client) (*TransferInCheckMailTokenResponse, error) {
	resp := &TransferInCheckMailTokenResponse{}
	err := client.Request("domain", "TransferInCheckMailToken", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TransferInCheckMailTokenResponse struct {
	RequestId   goaliyun.String
	SuccessList TransferInCheckMailTokenSuccessListList
	FailList    TransferInCheckMailTokenFailListList
}

type TransferInCheckMailTokenSuccessListList []goaliyun.String

func (list *TransferInCheckMailTokenSuccessListList) UnmarshalJSON(data []byte) error {
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

type TransferInCheckMailTokenFailListList []goaliyun.String

func (list *TransferInCheckMailTokenFailListList) UnmarshalJSON(data []byte) error {
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
