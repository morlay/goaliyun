package cms

import (
	"github.com/morlay/goaliyun"
)

type CreateMyGroupsRequest struct {
	ContactGroups string `position:"Query" name:"ContactGroups"`
	Options       string `position:"Query" name:"Options"`
	ServiceId     int64  `position:"Query" name:"ServiceId"`
	Type          string `position:"Query" name:"Type"`
	GroupName     string `position:"Query" name:"GroupName"`
	BindUrl       string `position:"Query" name:"BindUrl"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *CreateMyGroupsRequest) Invoke(client goaliyun.Client) (*CreateMyGroupsResponse, error) {
	resp := &CreateMyGroupsResponse{}
	err := client.Request("cms", "CreateMyGroups", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateMyGroupsResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	GroupId      goaliyun.Integer
}
