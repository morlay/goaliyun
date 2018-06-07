package ccc

import (
	"github.com/morlay/goaliyun"
)

type DeleteSkillGroupRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteSkillGroupRequest) Invoke(client goaliyun.Client) (*DeleteSkillGroupResponse, error) {
	resp := &DeleteSkillGroupResponse{}
	err := client.Request("ccc", "DeleteSkillGroup", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteSkillGroupResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}
