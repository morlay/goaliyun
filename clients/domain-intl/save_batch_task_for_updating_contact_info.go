package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveBatchTaskForUpdatingContactInfoRequest struct {
	ContactType         string                                             `position:"Query" name:"ContactType"`
	UserClientIp        string                                             `position:"Query" name:"UserClientIp"`
	RegistrantProfileId int64                                              `position:"Query" name:"RegistrantProfileId"`
	DomainNames         *SaveBatchTaskForUpdatingContactInfoDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	AddTransferLock     string                                             `position:"Query" name:"AddTransferLock"`
	Lang                string                                             `position:"Query" name:"Lang"`
	RegionId            string                                             `position:"Query" name:"RegionId"`
}

func (req *SaveBatchTaskForUpdatingContactInfoRequest) Invoke(client goaliyun.Client) (*SaveBatchTaskForUpdatingContactInfoResponse, error) {
	resp := &SaveBatchTaskForUpdatingContactInfoResponse{}
	err := client.Request("domain-intl", "SaveBatchTaskForUpdatingContactInfo", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchTaskForUpdatingContactInfoResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveBatchTaskForUpdatingContactInfoDomainNameList []string

func (list *SaveBatchTaskForUpdatingContactInfoDomainNameList) UnmarshalJSON(data []byte) error {
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
