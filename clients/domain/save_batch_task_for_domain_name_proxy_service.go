package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveBatchTaskForDomainNameProxyServiceRequest struct {
	UserClientIp string                                                `position:"Query" name:"UserClientIp"`
	DomainNames  *SaveBatchTaskForDomainNameProxyServiceDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang         string                                                `position:"Query" name:"Lang"`
	Status       string                                                `position:"Query" name:"Status"`
	RegionId     string                                                `position:"Query" name:"RegionId"`
}

func (req *SaveBatchTaskForDomainNameProxyServiceRequest) Invoke(client goaliyun.Client) (*SaveBatchTaskForDomainNameProxyServiceResponse, error) {
	resp := &SaveBatchTaskForDomainNameProxyServiceResponse{}
	err := client.Request("domain", "SaveBatchTaskForDomainNameProxyService", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchTaskForDomainNameProxyServiceResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveBatchTaskForDomainNameProxyServiceDomainNameList []string

func (list *SaveBatchTaskForDomainNameProxyServiceDomainNameList) UnmarshalJSON(data []byte) error {
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
