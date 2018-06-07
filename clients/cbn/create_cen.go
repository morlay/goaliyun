package cbn

import (
	"github.com/morlay/goaliyun"
)

type CreateCenRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateCenRequest) Invoke(client goaliyun.Client) (*CreateCenResponse, error) {
	resp := &CreateCenResponse{}
	err := client.Request("cbn", "CreateCen", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateCenResponse struct {
	RequestId goaliyun.String
	CenId     goaliyun.String
}
