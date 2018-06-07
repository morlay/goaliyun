package afs

import (
	"github.com/morlay/goaliyun"
)

type DescribeConfigNameRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeConfigNameRequest) Invoke(client goaliyun.Client) (*DescribeConfigNameResponse, error) {
	resp := &DescribeConfigNameResponse{}
	err := client.Request("afs", "DescribeConfigName", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeConfigNameResponse struct {
	RequestId   goaliyun.String
	HasConfig   bool
	ConfigNames goaliyun.String
	BizCode     goaliyun.String
}
