package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVSwitchesRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVSwitchesRequest) Invoke(client goaliyun.Client) (*DescribeVSwitchesResponse, error) {
	resp := &DescribeVSwitchesResponse{}
	err := client.Request("ecs", "DescribeVSwitches", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVSwitchesResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	VSwitches  DescribeVSwitchesVSwitchList
}

type DescribeVSwitchesVSwitch struct {
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
}

type DescribeVSwitchesVSwitchList []DescribeVSwitchesVSwitch

func (list *DescribeVSwitchesVSwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVSwitchesVSwitch)
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
