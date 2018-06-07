package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRdsVSwitchsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRdsVSwitchsRequest) Invoke(client goaliyun.Client) (*DescribeRdsVSwitchsResponse, error) {
	resp := &DescribeRdsVSwitchsResponse{}
	err := client.Request("r-kvstore", "DescribeRdsVSwitchs", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRdsVSwitchsResponse struct {
	RequestId goaliyun.String
	VSwitches DescribeRdsVSwitchsVSwitches
}

type DescribeRdsVSwitchsVSwitches struct {
	VSwitch DescribeRdsVSwitchsVSwitchItemList
}

type DescribeRdsVSwitchsVSwitchItem struct {
	VSwitchId   goaliyun.String
	VSwitchName goaliyun.String
	IzNo        goaliyun.String
	Bid         goaliyun.String
	AliUid      goaliyun.String
	RegionNo    goaliyun.String
	CidrBlock   goaliyun.String
	IsDefault   bool
	Status      goaliyun.String
	GmtCreate   goaliyun.String
	GmtModified goaliyun.String
}

type DescribeRdsVSwitchsVSwitchItemList []DescribeRdsVSwitchsVSwitchItem

func (list *DescribeRdsVSwitchsVSwitchItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRdsVSwitchsVSwitchItem)
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
