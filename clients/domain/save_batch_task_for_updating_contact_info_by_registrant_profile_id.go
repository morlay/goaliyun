package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest struct {
	ContactType           string                                                                  `position:"Query" name:"ContactType"`
	UserClientIp          string                                                                  `position:"Query" name:"UserClientIp"`
	RegistrantProfileId   int64                                                                   `position:"Query" name:"RegistrantProfileId"`
	DomainNames           *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	TransferOutProhibited string                                                                  `position:"Query" name:"TransferOutProhibited"`
	Lang                  string                                                                  `position:"Query" name:"Lang"`
	RegionId              string                                                                  `position:"Query" name:"RegionId"`
}

func (req *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) Invoke(client goaliyun.Client) (*SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse, error) {
	resp := &SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse{}
	err := client.Request("domain", "SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdDomainNameList []string

func (list *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdDomainNameList) UnmarshalJSON(data []byte) error {
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
