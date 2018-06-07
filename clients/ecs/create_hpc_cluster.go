package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateHpcClusterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Description          string `position:"Query" name:"Description"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateHpcClusterRequest) Invoke(client goaliyun.Client) (*CreateHpcClusterResponse, error) {
	resp := &CreateHpcClusterResponse{}
	err := client.Request("ecs", "CreateHpcCluster", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateHpcClusterResponse struct {
	RequestId    goaliyun.String
	HpcClusterId goaliyun.String
}
