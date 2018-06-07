package nas

import (
	"github.com/morlay/goaliyun"
)

type CreateAccessGroupRequest struct {
	Description     string `position:"Query" name:"Description"`
	AccessGroupType string `position:"Query" name:"AccessGroupType"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateAccessGroupRequest) Invoke(client goaliyun.Client) (*CreateAccessGroupResponse, error) {
	resp := &CreateAccessGroupResponse{}
	err := client.Request("nas", "CreateAccessGroup", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAccessGroupResponse struct {
	RequestId       goaliyun.String
	AccessGroupName goaliyun.String
}
