package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateDeploymentSetRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	DeploymentSetName    string `position:"Query" name:"DeploymentSetName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Granularity          string `position:"Query" name:"Granularity"`
	Domain               string `position:"Query" name:"Domain"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	Strategy             string `position:"Query" name:"Strategy"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateDeploymentSetRequest) Invoke(client goaliyun.Client) (*CreateDeploymentSetResponse, error) {
	resp := &CreateDeploymentSetResponse{}
	err := client.Request("ecs", "CreateDeploymentSet", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDeploymentSetResponse struct {
	RequestId       goaliyun.String
	DeploymentSetId goaliyun.String
}
