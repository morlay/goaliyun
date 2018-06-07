package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRdsVpcsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRdsVpcsRequest) Invoke(client goaliyun.Client) (*DescribeRdsVpcsResponse, error) {
	resp := &DescribeRdsVpcsResponse{}
	err := client.Request("r-kvstore", "DescribeRdsVpcs", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRdsVpcsResponse struct {
	RequestId goaliyun.String
	Vpcs      DescribeRdsVpcsVpcs
}

type DescribeRdsVpcsVpcs struct {
	Vpc DescribeRdsVpcsVpcItemList
}

type DescribeRdsVpcsVpcItem struct {
	VpcId       goaliyun.String
	VpcName     goaliyun.String
	Bid         goaliyun.String
	AliUid      goaliyun.String
	RegionNo    goaliyun.String
	CidrBlock   goaliyun.String
	IsDefault   bool
	Status      goaliyun.String
	GmtCreate   goaliyun.String
	GmtModified goaliyun.String
	VSwitchs    DescribeRdsVpcsVSwitchList
}

type DescribeRdsVpcsVSwitch struct {
	VSwitchId   goaliyun.String
	VSwitchName goaliyun.String
	IzNo        goaliyun.String
	CidrBlock   goaliyun.String
	IsDefault   bool
	Status      goaliyun.String
	GmtCreate   goaliyun.String
	GmtModified goaliyun.String
}

type DescribeRdsVpcsVpcItemList []DescribeRdsVpcsVpcItem

func (list *DescribeRdsVpcsVpcItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRdsVpcsVpcItem)
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

type DescribeRdsVpcsVSwitchList []DescribeRdsVpcsVSwitch

func (list *DescribeRdsVpcsVSwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRdsVpcsVSwitch)
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
