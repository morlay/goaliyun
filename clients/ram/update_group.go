package ram

import (
	"github.com/morlay/goaliyun"
)

type UpdateGroupRequest struct {
	NewGroupName string `position:"Query" name:"NewGroupName"`
	NewComments  string `position:"Query" name:"NewComments"`
	GroupName    string `position:"Query" name:"GroupName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *UpdateGroupRequest) Invoke(client goaliyun.Client) (*UpdateGroupResponse, error) {
	resp := &UpdateGroupResponse{}
	err := client.Request("ram", "UpdateGroup", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateGroupResponse struct {
	RequestId goaliyun.String
	Group     UpdateGroupGroup
}

type UpdateGroupGroup struct {
	GroupName  goaliyun.String
	Comments   goaliyun.String
	CreateDate goaliyun.String
	UpdateDate goaliyun.String
}
