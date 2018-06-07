package nas

import (
	"github.com/morlay/goaliyun"
)

type DeleteAccessGroupRequest struct {
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteAccessGroupRequest) Invoke(client goaliyun.Client) (*DeleteAccessGroupResponse, error) {
	resp := &DeleteAccessGroupResponse{}
	err := client.Request("nas", "DeleteAccessGroup", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAccessGroupResponse struct {
	RequestId goaliyun.String
}
