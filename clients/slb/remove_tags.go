package slb

import (
	"github.com/morlay/goaliyun"
)

type RemoveTagsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RemoveTagsRequest) Invoke(client goaliyun.Client) (*RemoveTagsResponse, error) {
	resp := &RemoveTagsResponse{}
	err := client.Request("slb", "RemoveTags", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveTagsResponse struct {
	RequestId goaliyun.String
}
