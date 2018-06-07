package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVSwitchAttributesRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVSwitchAttributesRequest) Invoke(client goaliyun.Client) (*DescribeVSwitchAttributesResponse, error) {
	resp := &DescribeVSwitchAttributesResponse{}
	err := client.Request("vpc", "DescribeVSwitchAttributes", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVSwitchAttributesResponse struct {
	RequestId               goaliyun.String
	VSwitchId               goaliyun.String
	VpcId                   goaliyun.String
	Status                  goaliyun.String
	CidrBlock               goaliyun.String
	ZoneId                  goaliyun.String
	AvailableIpAddressCount goaliyun.Integer
	Description             goaliyun.String
	VSwitchName             goaliyun.String
	CreationTime            goaliyun.String
	IsDefault               bool
	CloudResources          DescribeVSwitchAttributesCloudResourceSetTypeList
}

type DescribeVSwitchAttributesCloudResourceSetType struct {
	ResourceType  goaliyun.String
	ResourceCount goaliyun.Integer
}

type DescribeVSwitchAttributesCloudResourceSetTypeList []DescribeVSwitchAttributesCloudResourceSetType

func (list *DescribeVSwitchAttributesCloudResourceSetTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVSwitchAttributesCloudResourceSetType)
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
