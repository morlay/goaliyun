package alidns

import (
	"github.com/morlay/goaliyun"
)

type DeleteBatchDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Records      string `position:"Query" name:"Records"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteBatchDomainRecordsRequest) Invoke(client goaliyun.Client) (*DeleteBatchDomainRecordsResponse, error) {
	resp := &DeleteBatchDomainRecordsResponse{}
	err := client.Request("alidns", "DeleteBatchDomainRecords", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteBatchDomainRecordsResponse struct {
	RequestId goaliyun.String
	TraceId   goaliyun.String
}
