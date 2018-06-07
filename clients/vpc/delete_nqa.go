package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteNqaRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NqaId                string `position:"Query" name:"NqaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteNqaRequest) Invoke(client goaliyun.Client) (*DeleteNqaResponse, error) {
	resp := &DeleteNqaResponse{}
	err := client.Request("vpc", "DeleteNqa", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteNqaResponse struct {
	RequestId goaliyun.String
}
