package alidns

import (
	"github.com/morlay/goaliyun"
)

type UpdateDomainRecordRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
	RR           string `position:"Query" name:"RR"`
	Type         string `position:"Query" name:"Type"`
	Value        string `position:"Query" name:"Value"`
	TTL          int64  `position:"Query" name:"TTL"`
	Priority     int64  `position:"Query" name:"Priority"`
	Line         string `position:"Query" name:"Line"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *UpdateDomainRecordRequest) Invoke(client goaliyun.Client) (*UpdateDomainRecordResponse, error) {
	resp := &UpdateDomainRecordResponse{}
	err := client.Request("alidns", "UpdateDomainRecord", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateDomainRecordResponse struct {
	RequestId goaliyun.String
	RecordId  goaliyun.String
}
