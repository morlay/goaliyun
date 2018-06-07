package slb

import (
	"github.com/morlay/goaliyun"
)

type DescribeListenerAccessControlAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int64  `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeListenerAccessControlAttributeRequest) Invoke(client goaliyun.Client) (*DescribeListenerAccessControlAttributeResponse, error) {
	resp := &DescribeListenerAccessControlAttributeResponse{}
	err := client.Request("slb", "DescribeListenerAccessControlAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeListenerAccessControlAttributeResponse struct {
	RequestId           goaliyun.String
	AccessControlStatus goaliyun.String
	SourceItems         goaliyun.String
}
