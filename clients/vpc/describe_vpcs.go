package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVpcsRequest struct {
	VpcName              string `position:"Query" name:"VpcName"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVpcsRequest) Invoke(client goaliyun.Client) (*DescribeVpcsResponse, error) {
	resp := &DescribeVpcsResponse{}
	err := client.Request("vpc", "DescribeVpcs", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVpcsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Vpcs       DescribeVpcsVpcList
}

type DescribeVpcsVpc struct {
	VpcId           goaliyun.String
	RegionId        goaliyun.String
	Status          goaliyun.String
	VpcName         goaliyun.String
	CreationTime    goaliyun.String
	CidrBlock       goaliyun.String
	VRouterId       goaliyun.String
	Description     goaliyun.String
	IsDefault       bool
	ResourceGroupId goaliyun.String
	VSwitchIds      DescribeVpcsVSwitchIdList
	UserCidrs       DescribeVpcsUserCidrList
	NatGatewayIds   DescribeVpcsNatGatewayIdList
	RouterTableIds  DescribeVpcsRouterTableIdList
}

type DescribeVpcsVpcList []DescribeVpcsVpc

func (list *DescribeVpcsVpcList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcsVpc)
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

type DescribeVpcsVSwitchIdList []goaliyun.String

func (list *DescribeVpcsVSwitchIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type DescribeVpcsUserCidrList []goaliyun.String

func (list *DescribeVpcsUserCidrList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type DescribeVpcsNatGatewayIdList []goaliyun.String

func (list *DescribeVpcsNatGatewayIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type DescribeVpcsRouterTableIdList []goaliyun.String

func (list *DescribeVpcsRouterTableIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
