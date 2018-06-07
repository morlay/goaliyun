package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyVpcAttributeRequest struct {
	VpcName              string `position:"Query" name:"VpcName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyVpcAttributeRequest) Invoke(client goaliyun.Client) (*ModifyVpcAttributeResponse, error) {
	resp := &ModifyVpcAttributeResponse{}
	err := client.Request("vpc", "ModifyVpcAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyVpcAttributeResponse struct {
	RequestId goaliyun.String
}
