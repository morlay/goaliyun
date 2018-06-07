package vpc

import (
	"github.com/morlay/goaliyun"
)

type DescribeNetworkQuotasRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Product              string `position:"Query" name:"Product"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeNetworkQuotasRequest) Invoke(client goaliyun.Client) (*DescribeNetworkQuotasResponse, error) {
	resp := &DescribeNetworkQuotasResponse{}
	err := client.Request("vpc", "DescribeNetworkQuotas", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeNetworkQuotasResponse struct {
	RequestId goaliyun.String
	Product   goaliyun.String
	RegionId  goaliyun.String
	Quota     goaliyun.String
}
