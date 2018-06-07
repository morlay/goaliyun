package market

import (
	"github.com/morlay/goaliyun"
)

type DescribeInstanceRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceRequest) Invoke(client goaliyun.Client) (*DescribeInstanceResponse, error) {
	resp := &DescribeInstanceResponse{}
	err := client.Request("market", "DescribeInstance", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceResponse struct {
	InstanceId     goaliyun.Integer
	OrderId        goaliyun.Integer
	SupplierName   goaliyun.String
	ProductCode    goaliyun.String
	ProductSkuCode goaliyun.String
	ProductName    goaliyun.String
	ProductType    goaliyun.String
	Status         goaliyun.String
	BeganOn        goaliyun.Integer
	EndOn          goaliyun.Integer
	CreatedOn      goaliyun.Integer
	ExtendJson     goaliyun.String
	HostJson       goaliyun.String
	AppJson        goaliyun.String
}
