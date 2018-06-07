package slb

import (
	"github.com/morlay/goaliyun"
)

type SetListenerAccessControlStatusRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int64  `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AccessControlStatus  string `position:"Query" name:"AccessControlStatus"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetListenerAccessControlStatusRequest) Invoke(client goaliyun.Client) (*SetListenerAccessControlStatusResponse, error) {
	resp := &SetListenerAccessControlStatusResponse{}
	err := client.Request("slb", "SetListenerAccessControlStatus", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetListenerAccessControlStatusResponse struct {
	RequestId goaliyun.String
}
