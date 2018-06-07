package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteVpcRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteVpcRequest) Invoke(client goaliyun.Client) (*DeleteVpcResponse, error) {
	resp := &DeleteVpcResponse{}
	err := client.Request("ecs", "DeleteVpc", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteVpcResponse struct {
	RequestId goaliyun.String
}
