package alidns

import (
	"github.com/morlay/goaliyun"
)

type AddBatchDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Records      string `position:"Query" name:"Records"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *AddBatchDomainRecordsRequest) Invoke(client goaliyun.Client) (*AddBatchDomainRecordsResponse, error) {
	resp := &AddBatchDomainRecordsResponse{}
	err := client.Request("alidns", "AddBatchDomainRecords", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddBatchDomainRecordsResponse struct {
	RequestId goaliyun.String
	TraceId   goaliyun.String
}
