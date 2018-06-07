package cms

import (
	"github.com/morlay/goaliyun"
)

type UpdateMyGroupMembersRequest struct {
	Readers  string `position:"Query" name:"Readers"`
	GroupId  int64  `position:"Query" name:"GroupId"`
	Masters  string `position:"Query" name:"Masters"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateMyGroupMembersRequest) Invoke(client goaliyun.Client) (*UpdateMyGroupMembersResponse, error) {
	resp := &UpdateMyGroupMembersResponse{}
	err := client.Request("cms", "UpdateMyGroupMembers", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateMyGroupMembersResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
}
