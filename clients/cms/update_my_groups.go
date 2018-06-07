package cms

import (
	"github.com/morlay/goaliyun"
)

type UpdateMyGroupsRequest struct {
	ContactGroups string `position:"Query" name:"ContactGroups"`
	GroupId       string `position:"Query" name:"GroupId"`
	ServiceId     int64  `position:"Query" name:"ServiceId"`
	Type          string `position:"Query" name:"Type"`
	GroupName     string `position:"Query" name:"GroupName"`
	BindUrls      string `position:"Query" name:"BindUrls"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *UpdateMyGroupsRequest) Invoke(client goaliyun.Client) (*UpdateMyGroupsResponse, error) {
	resp := &UpdateMyGroupsResponse{}
	err := client.Request("cms", "UpdateMyGroups", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateMyGroupsResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
}
