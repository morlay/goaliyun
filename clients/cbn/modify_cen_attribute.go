package cbn

import (
	"github.com/morlay/goaliyun"
)

type ModifyCenAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyCenAttributeRequest) Invoke(client goaliyun.Client) (*ModifyCenAttributeResponse, error) {
	resp := &ModifyCenAttributeResponse{}
	err := client.Request("cbn", "ModifyCenAttribute", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCenAttributeResponse struct {
	RequestId goaliyun.String
}
