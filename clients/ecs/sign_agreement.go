package ecs

import (
	"github.com/morlay/goaliyun"
)

type SignAgreementRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AgreementType        string `position:"Query" name:"AgreementType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SignAgreementRequest) Invoke(client goaliyun.Client) (*SignAgreementResponse, error) {
	resp := &SignAgreementResponse{}
	err := client.Request("ecs", "SignAgreement", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SignAgreementResponse struct {
	RequestId goaliyun.String
}
