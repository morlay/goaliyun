package slb

import (
	"github.com/morlay/goaliyun"
)

type AddTagsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddTagsRequest) Invoke(client goaliyun.Client) (*AddTagsResponse, error) {
	resp := &AddTagsResponse{}
	err := client.Request("slb", "AddTags", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddTagsResponse struct {
	RequestId goaliyun.String
}
