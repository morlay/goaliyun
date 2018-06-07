package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateScalingGroupRequest struct {
	MultiAZPolicy        string                           `position:"Query" name:"MultiAZPolicy"`
	DBInstanceIds        string                           `position:"Query" name:"DBInstanceIds"`
	LoadBalancerIds      string                           `position:"Query" name:"LoadBalancerIds"`
	ResourceOwnerAccount string                           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName     string                           `position:"Query" name:"ScalingGroupName"`
	VSwitchIds           *CreateScalingGroupVSwitchIdList `position:"Query" type:"Repeated" name:"VSwitchId"`
	OwnerAccount         string                           `position:"Query" name:"OwnerAccount"`
	MinSize              int64                            `position:"Query" name:"MinSize"`
	OwnerId              int64                            `position:"Query" name:"OwnerId"`
	VSwitchId            string                           `position:"Query" name:"VSwitchId"`
	MaxSize              int64                            `position:"Query" name:"MaxSize"`
	DefaultCooldown      int64                            `position:"Query" name:"DefaultCooldown"`
	RemovalPolicy1       string                           `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2       string                           `position:"Query" name:"RemovalPolicy.2"`
	RegionId             string                           `position:"Query" name:"RegionId"`
}

func (req *CreateScalingGroupRequest) Invoke(client goaliyun.Client) (*CreateScalingGroupResponse, error) {
	resp := &CreateScalingGroupResponse{}
	err := client.Request("ess", "CreateScalingGroup", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateScalingGroupResponse struct {
	ScalingGroupId goaliyun.String
	RequestId      goaliyun.String
}

type CreateScalingGroupVSwitchIdList []string

func (list *CreateScalingGroupVSwitchIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
