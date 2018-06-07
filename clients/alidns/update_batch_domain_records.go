package alidns

import (
	"github.com/morlay/goaliyun"
)

type UpdateBatchDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Records      string `position:"Query" name:"Records"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *UpdateBatchDomainRecordsRequest) Invoke(client goaliyun.Client) (*UpdateBatchDomainRecordsResponse, error) {
	resp := &UpdateBatchDomainRecordsResponse{}
	err := client.Request("alidns", "UpdateBatchDomainRecords", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateBatchDomainRecordsResponse struct {
	RequestId goaliyun.String
	TraceId   goaliyun.String
}
