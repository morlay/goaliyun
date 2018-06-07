package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateVSwitchRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpcId                string `position:"Query" name:"VpcId"`
	VSwitchName          string `position:"Query" name:"VSwitchName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CidrBlock            string `position:"Query" name:"CidrBlock"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateVSwitchRequest) Invoke(client goaliyun.Client) (*CreateVSwitchResponse, error) {
	resp := &CreateVSwitchResponse{}
	err := client.Request("vpc", "CreateVSwitch", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateVSwitchResponse struct {
	RequestId goaliyun.String
	VSwitchId goaliyun.String
}
