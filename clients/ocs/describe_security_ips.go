package ocs

import (
	"github.com/morlay/goaliyun"
)

type DescribeSecurityIpsRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSecurityIpsRequest) Invoke(client goaliyun.Client) (*DescribeSecurityIpsResponse, error) {
	resp := &DescribeSecurityIpsResponse{}
	err := client.Request("ocs", "DescribeSecurityIps", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSecurityIpsResponse struct {
	RequestId                         goaliyun.String
	DescribeOcsSecurityIpsResponseDTO DescribeSecurityIpsDescribeOcsSecurityIpsResponseDTO
}

type DescribeSecurityIpsDescribeOcsSecurityIpsResponseDTO struct {
	SecurityIps goaliyun.String
}
