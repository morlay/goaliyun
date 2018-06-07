package ram

import (
	"github.com/morlay/goaliyun"
)

type CreateGroupRequest struct {
	Comments  string `position:"Query" name:"Comments"`
	GroupName string `position:"Query" name:"GroupName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *CreateGroupRequest) Invoke(client goaliyun.Client) (*CreateGroupResponse, error) {
	resp := &CreateGroupResponse{}
	err := client.Request("ram", "CreateGroup", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateGroupResponse struct {
	RequestId goaliyun.String
	Group     CreateGroupGroup
}

type CreateGroupGroup struct {
	GroupName  goaliyun.String
	Comments   goaliyun.String
	CreateDate goaliyun.String
}
