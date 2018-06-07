package slb

import (
	"github.com/morlay/goaliyun"
)

type RemoveListenerWhiteListItemRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int64  `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	SourceItems          string `position:"Query" name:"SourceItems"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RemoveListenerWhiteListItemRequest) Invoke(client goaliyun.Client) (*RemoveListenerWhiteListItemResponse, error) {
	resp := &RemoveListenerWhiteListItemResponse{}
	err := client.Request("slb", "RemoveListenerWhiteListItem", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveListenerWhiteListItemResponse struct {
	RequestId goaliyun.String
}
