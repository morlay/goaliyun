package domain

import (
	"github.com/morlay/goaliyun"
)

type SaveBatchDomainRemarkRequest struct {
	InstanceIds string `position:"Query" name:"InstanceIds"`
	Remark      string `position:"Query" name:"Remark"`
	Lang        string `position:"Query" name:"Lang"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *SaveBatchDomainRemarkRequest) Invoke(client goaliyun.Client) (*SaveBatchDomainRemarkResponse, error) {
	resp := &SaveBatchDomainRemarkResponse{}
	err := client.Request("domain", "SaveBatchDomainRemark", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchDomainRemarkResponse struct {
	RequestId goaliyun.String
}
