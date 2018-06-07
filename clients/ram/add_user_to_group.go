package ram

import (
	"github.com/morlay/goaliyun"
)

type AddUserToGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
	UserName  string `position:"Query" name:"UserName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *AddUserToGroupRequest) Invoke(client goaliyun.Client) (*AddUserToGroupResponse, error) {
	resp := &AddUserToGroupResponse{}
	err := client.Request("ram", "AddUserToGroup", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddUserToGroupResponse struct {
	RequestId goaliyun.String
}
