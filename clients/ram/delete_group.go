package ram

import (
	"github.com/morlay/goaliyun"
)

type DeleteGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DeleteGroupRequest) Invoke(client goaliyun.Client) (*DeleteGroupResponse, error) {
	resp := &DeleteGroupResponse{}
	err := client.Request("ram", "DeleteGroup", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteGroupResponse struct {
	RequestId goaliyun.String
}
