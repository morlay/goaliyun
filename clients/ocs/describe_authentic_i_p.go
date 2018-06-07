package ocs

import (
	"github.com/morlay/goaliyun"
)

type DescribeAuthenticIPRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAuthenticIPRequest) Invoke(client goaliyun.Client) (*DescribeAuthenticIPResponse, error) {
	resp := &DescribeAuthenticIPResponse{}
	err := client.Request("ocs", "DescribeAuthenticIP", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAuthenticIPResponse struct {
	RequestId              goaliyun.String
	GetAuthenticIpResponse DescribeAuthenticIPGetAuthenticIpResponse
}

type DescribeAuthenticIPGetAuthenticIpResponse struct {
	AuthenticIPs goaliyun.String
}
