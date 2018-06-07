package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteDeploymentSetRequest struct {
	DeploymentSetId      string `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteDeploymentSetRequest) Invoke(client goaliyun.Client) (*DeleteDeploymentSetResponse, error) {
	resp := &DeleteDeploymentSetResponse{}
	err := client.Request("ecs", "DeleteDeploymentSet", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDeploymentSetResponse struct {
	RequestId goaliyun.String
}
