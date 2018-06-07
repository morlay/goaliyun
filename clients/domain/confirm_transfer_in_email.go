package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ConfirmTransferInEmailRequest struct {
	UserClientIp string                                `position:"Query" name:"UserClientIp"`
	DomainNames  *ConfirmTransferInEmailDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang         string                                `position:"Query" name:"Lang"`
	Email        string                                `position:"Query" name:"Email"`
	RegionId     string                                `position:"Query" name:"RegionId"`
}

func (req *ConfirmTransferInEmailRequest) Invoke(client goaliyun.Client) (*ConfirmTransferInEmailResponse, error) {
	resp := &ConfirmTransferInEmailResponse{}
	err := client.Request("domain", "ConfirmTransferInEmail", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ConfirmTransferInEmailResponse struct {
	RequestId   goaliyun.String
	SuccessList ConfirmTransferInEmailSuccessListList
	FailList    ConfirmTransferInEmailFailListList
}

type ConfirmTransferInEmailDomainNameList []string

func (list *ConfirmTransferInEmailDomainNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ConfirmTransferInEmailSuccessListList []goaliyun.String

func (list *ConfirmTransferInEmailSuccessListList) UnmarshalJSON(data []byte) error {
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

type ConfirmTransferInEmailFailListList []goaliyun.String

func (list *ConfirmTransferInEmailFailListList) UnmarshalJSON(data []byte) error {
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
