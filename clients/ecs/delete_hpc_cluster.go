package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteHpcClusterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	HpcClusterId         string `position:"Query" name:"HpcClusterId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteHpcClusterRequest) Invoke(client goaliyun.Client) (*DeleteHpcClusterResponse, error) {
	resp := &DeleteHpcClusterResponse{}
	err := client.Request("ecs", "DeleteHpcCluster", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteHpcClusterResponse struct {
	RequestId goaliyun.String
}
