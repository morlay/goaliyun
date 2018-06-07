package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveBatchTaskForUpdateProhibitionLockRequest struct {
	UserClientIp string                                               `position:"Query" name:"UserClientIp"`
	DomainNames  *SaveBatchTaskForUpdateProhibitionLockDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang         string                                               `position:"Query" name:"Lang"`
	Status       string                                               `position:"Query" name:"Status"`
	RegionId     string                                               `position:"Query" name:"RegionId"`
}

func (req *SaveBatchTaskForUpdateProhibitionLockRequest) Invoke(client goaliyun.Client) (*SaveBatchTaskForUpdateProhibitionLockResponse, error) {
	resp := &SaveBatchTaskForUpdateProhibitionLockResponse{}
	err := client.Request("domain", "SaveBatchTaskForUpdateProhibitionLock", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchTaskForUpdateProhibitionLockResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveBatchTaskForUpdateProhibitionLockDomainNameList []string

func (list *SaveBatchTaskForUpdateProhibitionLockDomainNameList) UnmarshalJSON(data []byte) error {
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
