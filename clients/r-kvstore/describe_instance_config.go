package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type DescribeInstanceConfigRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceConfigRequest) Invoke(client goaliyun.Client) (*DescribeInstanceConfigResponse, error) {
	resp := &DescribeInstanceConfigResponse{}
	err := client.Request("r-kvstore", "DescribeInstanceConfig", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceConfigResponse struct {
	RequestId goaliyun.String
	Config    goaliyun.String
}
