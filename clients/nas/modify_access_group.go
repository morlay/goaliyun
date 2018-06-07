package nas

import (
	"github.com/morlay/goaliyun"
)

type ModifyAccessGroupRequest struct {
	Description     string `position:"Query" name:"Description"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyAccessGroupRequest) Invoke(client goaliyun.Client) (*ModifyAccessGroupResponse, error) {
	resp := &ModifyAccessGroupResponse{}
	err := client.Request("nas", "ModifyAccessGroup", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyAccessGroupResponse struct {
	RequestId goaliyun.String
}
