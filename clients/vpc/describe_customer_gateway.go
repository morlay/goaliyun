package vpc

import (
	"github.com/morlay/goaliyun"
)

type DescribeCustomerGatewayRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCustomerGatewayRequest) Invoke(client goaliyun.Client) (*DescribeCustomerGatewayResponse, error) {
	resp := &DescribeCustomerGatewayResponse{}
	err := client.Request("vpc", "DescribeCustomerGateway", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCustomerGatewayResponse struct {
	RequestId         goaliyun.String
	CustomerGatewayId goaliyun.String
	IpAddress         goaliyun.String
	Name              goaliyun.String
	Description       goaliyun.String
	CreateTime        goaliyun.Integer
}
