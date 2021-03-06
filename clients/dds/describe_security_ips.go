package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSecurityIpsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSecurityIpsRequest) Invoke(client goaliyun.Client) (*DescribeSecurityIpsResponse, error) {
	resp := &DescribeSecurityIpsResponse{}
	err := client.Request("dds", "DescribeSecurityIps", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSecurityIpsResponse struct {
	RequestId        goaliyun.String
	SecurityIps      goaliyun.String
	SecurityIpGroups DescribeSecurityIpsSecurityIpGroupList
}

type DescribeSecurityIpsSecurityIpGroup struct {
	SecurityIpGroupName      goaliyun.String
	SecurityIpGroupAttribute goaliyun.String
	SecurityIpList           goaliyun.String
}

type DescribeSecurityIpsSecurityIpGroupList []DescribeSecurityIpsSecurityIpGroup

func (list *DescribeSecurityIpsSecurityIpGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityIpsSecurityIpGroup)
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
