package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateNqaRequest struct {
	DestinationIp        string `position:"Query" name:"DestinationIp"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateNqaRequest) Invoke(client goaliyun.Client) (*CreateNqaResponse, error) {
	resp := &CreateNqaResponse{}
	err := client.Request("vpc", "CreateNqa", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateNqaResponse struct {
	RequestId goaliyun.String
	NqaId     goaliyun.String
}
