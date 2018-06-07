package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteVSwitchRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteVSwitchRequest) Invoke(client goaliyun.Client) (*DeleteVSwitchResponse, error) {
	resp := &DeleteVSwitchResponse{}
	err := client.Request("vpc", "DeleteVSwitch", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteVSwitchResponse struct {
	RequestId goaliyun.String
}
