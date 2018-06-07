package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyNqaRequest struct {
	DestinationIp        string `position:"Query" name:"DestinationIp"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NqaId                string `position:"Query" name:"NqaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyNqaRequest) Invoke(client goaliyun.Client) (*ModifyNqaResponse, error) {
	resp := &ModifyNqaResponse{}
	err := client.Request("vpc", "ModifyNqa", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyNqaResponse struct {
	RequestId goaliyun.String
}
