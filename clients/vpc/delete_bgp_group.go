package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteBgpGroupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteBgpGroupRequest) Invoke(client goaliyun.Client) (*DeleteBgpGroupResponse, error) {
	resp := &DeleteBgpGroupResponse{}
	err := client.Request("vpc", "DeleteBgpGroup", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteBgpGroupResponse struct {
	RequestId goaliyun.String
}
