package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDeploymentSetsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	DeploymentSetName    string `position:"Query" name:"DeploymentSetName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	DeploymentSetIds     string `position:"Query" name:"DeploymentSetIds"`
	Granularity          string `position:"Query" name:"Granularity"`
	Domain               string `position:"Query" name:"Domain"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	Strategy             string `position:"Query" name:"Strategy"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDeploymentSetsRequest) Invoke(client goaliyun.Client) (*DescribeDeploymentSetsResponse, error) {
	resp := &DescribeDeploymentSetsResponse{}
	err := client.Request("ecs", "DescribeDeploymentSets", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDeploymentSetsResponse struct {
	RequestId      goaliyun.String
	RegionId       goaliyun.String
	TotalCount     goaliyun.Integer
	PageNumber     goaliyun.Integer
	PageSize       goaliyun.Integer
	DeploymentSets DescribeDeploymentSetsDeploymentSetList
}

type DescribeDeploymentSetsDeploymentSet struct {
	DeploymentSetId          goaliyun.String
	DeploymentSetDescription goaliyun.String
	DeploymentSetName        goaliyun.String
	Strategy                 goaliyun.String
	Domain                   goaliyun.String
	Granularity              goaliyun.String
	InstanceAmount           goaliyun.Integer
	CreationTime             goaliyun.String
}

type DescribeDeploymentSetsDeploymentSetList []DescribeDeploymentSetsDeploymentSet

func (list *DescribeDeploymentSetsDeploymentSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetsDeploymentSet)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
