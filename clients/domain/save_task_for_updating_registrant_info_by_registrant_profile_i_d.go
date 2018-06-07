package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest struct {
	UserClientIp          string                                                                `position:"Query" name:"UserClientIp"`
	RegistrantProfileId   int64                                                                 `position:"Query" name:"RegistrantProfileId"`
	DomainNames           *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	TransferOutProhibited string                                                                `position:"Query" name:"TransferOutProhibited"`
	Lang                  string                                                                `position:"Query" name:"Lang"`
	RegionId              string                                                                `position:"Query" name:"RegionId"`
}

func (req *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) Invoke(client goaliyun.Client) (*SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse, error) {
	resp := &SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse{}
	err := client.Request("domain", "SaveTaskForUpdatingRegistrantInfoByRegistrantProfileID", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDDomainNameList []string

func (list *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDDomainNameList) UnmarshalJSON(data []byte) error {
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
