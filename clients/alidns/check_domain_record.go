package alidns

import (
	"github.com/morlay/goaliyun"
)

type CheckDomainRecordRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RR           string `position:"Query" name:"RR"`
	Type         string `position:"Query" name:"Type"`
	Value        string `position:"Query" name:"Value"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *CheckDomainRecordRequest) Invoke(client goaliyun.Client) (*CheckDomainRecordResponse, error) {
	resp := &CheckDomainRecordResponse{}
	err := client.Request("alidns", "CheckDomainRecord", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckDomainRecordResponse struct {
	RequestId goaliyun.String
	IsExist   bool
}
