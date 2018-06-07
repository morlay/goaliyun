package alidns

import (
	"github.com/morlay/goaliyun"
)

type DeleteSubDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RR           string `position:"Query" name:"RR"`
	Type         string `position:"Query" name:"Type"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteSubDomainRecordsRequest) Invoke(client goaliyun.Client) (*DeleteSubDomainRecordsResponse, error) {
	resp := &DeleteSubDomainRecordsResponse{}
	err := client.Request("alidns", "DeleteSubDomainRecords", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteSubDomainRecordsResponse struct {
	RequestId  goaliyun.String
	RR         goaliyun.String
	TotalCount goaliyun.String
}
