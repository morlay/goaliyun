package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCustomerGatewaysRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCustomerGatewaysRequest) Invoke(client goaliyun.Client) (*DescribeCustomerGatewaysResponse, error) {
	resp := &DescribeCustomerGatewaysResponse{}
	err := client.Request("vpc", "DescribeCustomerGateways", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCustomerGatewaysResponse struct {
	RequestId        goaliyun.String
	TotalCount       goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageSize         goaliyun.Integer
	CustomerGateways DescribeCustomerGatewaysCustomerGatewayList
}

type DescribeCustomerGatewaysCustomerGateway struct {
	CustomerGatewayId goaliyun.String
	Name              goaliyun.String
	IpAddress         goaliyun.String
	Description       goaliyun.String
	CreateTime        goaliyun.Integer
}

type DescribeCustomerGatewaysCustomerGatewayList []DescribeCustomerGatewaysCustomerGateway

func (list *DescribeCustomerGatewaysCustomerGatewayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCustomerGatewaysCustomerGateway)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
