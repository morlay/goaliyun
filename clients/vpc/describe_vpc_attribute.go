package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVpcAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVpcAttributeRequest) Invoke(client goaliyun.Client) (*DescribeVpcAttributeResponse, error) {
	resp := &DescribeVpcAttributeResponse{}
	err := client.Request("vpc", "DescribeVpcAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVpcAttributeResponse struct {
	RequestId          goaliyun.String
	VpcId              goaliyun.String
	RegionId           goaliyun.String
	Status             goaliyun.String
	VpcName            goaliyun.String
	CreationTime       goaliyun.String
	CidrBlock          goaliyun.String
	VRouterId          goaliyun.String
	Description        goaliyun.String
	IsDefault          bool
	ClassicLinkEnabled bool
	ResourceGroupId    goaliyun.String
	AssociatedCens     DescribeVpcAttributeAssociatedCenList
	CloudResources     DescribeVpcAttributeCloudResourceSetTypeList
	VSwitchIds         DescribeVpcAttributeVSwitchIdList
	UserCidrs          DescribeVpcAttributeUserCidrList
}

type DescribeVpcAttributeAssociatedCen struct {
	CenId      goaliyun.String
	CenOwnerId goaliyun.Integer
	CenStatus  goaliyun.String
}

type DescribeVpcAttributeCloudResourceSetType struct {
	ResourceType  goaliyun.String
	ResourceCount goaliyun.Integer
}

type DescribeVpcAttributeAssociatedCenList []DescribeVpcAttributeAssociatedCen

func (list *DescribeVpcAttributeAssociatedCenList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcAttributeAssociatedCen)
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

type DescribeVpcAttributeCloudResourceSetTypeList []DescribeVpcAttributeCloudResourceSetType

func (list *DescribeVpcAttributeCloudResourceSetTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcAttributeCloudResourceSetType)
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

type DescribeVpcAttributeVSwitchIdList []goaliyun.String

func (list *DescribeVpcAttributeVSwitchIdList) UnmarshalJSON(data []byte) error {
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

type DescribeVpcAttributeUserCidrList []goaliyun.String

func (list *DescribeVpcAttributeUserCidrList) UnmarshalJSON(data []byte) error {
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
