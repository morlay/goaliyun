package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveBatchTaskForTransferProhibitionLockRequest struct {
	UserClientIp string                                                 `position:"Query" name:"UserClientIp"`
	DomainNames  *SaveBatchTaskForTransferProhibitionLockDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang         string                                                 `position:"Query" name:"Lang"`
	Status       string                                                 `position:"Query" name:"Status"`
	RegionId     string                                                 `position:"Query" name:"RegionId"`
}

func (req *SaveBatchTaskForTransferProhibitionLockRequest) Invoke(client goaliyun.Client) (*SaveBatchTaskForTransferProhibitionLockResponse, error) {
	resp := &SaveBatchTaskForTransferProhibitionLockResponse{}
	err := client.Request("domain-intl", "SaveBatchTaskForTransferProhibitionLock", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchTaskForTransferProhibitionLockResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveBatchTaskForTransferProhibitionLockDomainNameList []string

func (list *SaveBatchTaskForTransferProhibitionLockDomainNameList) UnmarshalJSON(data []byte) error {
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
