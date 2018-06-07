package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyVSwitchAttributeRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VSwitchName          string `position:"Query" name:"VSwitchName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyVSwitchAttributeRequest) Invoke(client goaliyun.Client) (*ModifyVSwitchAttributeResponse, error) {
	resp := &ModifyVSwitchAttributeResponse{}
	err := client.Request("vpc", "ModifyVSwitchAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyVSwitchAttributeResponse struct {
	RequestId goaliyun.String
}
