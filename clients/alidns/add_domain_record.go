package alidns

import (
	"github.com/morlay/goaliyun"
)

type AddDomainRecordRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RR           string `position:"Query" name:"RR"`
	Type         string `position:"Query" name:"Type"`
	Value        string `position:"Query" name:"Value"`
	TTL          int64  `position:"Query" name:"TTL"`
	Priority     int64  `position:"Query" name:"Priority"`
	Line         string `position:"Query" name:"Line"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *AddDomainRecordRequest) Invoke(client goaliyun.Client) (*AddDomainRecordResponse, error) {
	resp := &AddDomainRecordResponse{}
	err := client.Request("alidns", "AddDomainRecord", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddDomainRecordResponse struct {
	RequestId goaliyun.String
	RecordId  goaliyun.String
}
