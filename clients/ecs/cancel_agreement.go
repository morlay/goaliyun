package ecs

import (
	"github.com/morlay/goaliyun"
)

type CancelAgreementRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AgreementType        string `position:"Query" name:"AgreementType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CancelAgreementRequest) Invoke(client goaliyun.Client) (*CancelAgreementResponse, error) {
	resp := &CancelAgreementResponse{}
	err := client.Request("ecs", "CancelAgreement", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelAgreementResponse struct {
	RequestId goaliyun.String
}
