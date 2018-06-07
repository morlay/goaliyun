package ram

import (
	"github.com/morlay/goaliyun"
)

type RemoveUserFromGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
	UserName  string `position:"Query" name:"UserName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *RemoveUserFromGroupRequest) Invoke(client goaliyun.Client) (*RemoveUserFromGroupResponse, error) {
	resp := &RemoveUserFromGroupResponse{}
	err := client.Request("ram", "RemoveUserFromGroup", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveUserFromGroupResponse struct {
	RequestId goaliyun.String
}
