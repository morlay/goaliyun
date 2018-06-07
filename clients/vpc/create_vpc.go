package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateVpcRequest struct {
	VpcName              string `position:"Query" name:"VpcName"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CidrBlock            string `position:"Query" name:"CidrBlock"`
	Description          string `position:"Query" name:"Description"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateVpcRequest) Invoke(client goaliyun.Client) (*CreateVpcResponse, error) {
	resp := &CreateVpcResponse{}
	err := client.Request("vpc", "CreateVpc", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateVpcResponse struct {
	RequestId       goaliyun.String
	VpcId           goaliyun.String
	VRouterId       goaliyun.String
	RouteTableId    goaliyun.String
	ResourceGroupId goaliyun.String
}
