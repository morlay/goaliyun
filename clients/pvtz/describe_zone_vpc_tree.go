package pvtz

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeZoneVpcTreeRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeZoneVpcTreeRequest) Invoke(client goaliyun.Client) (*DescribeZoneVpcTreeResponse, error) {
	resp := &DescribeZoneVpcTreeResponse{}
	err := client.Request("pvtz", "DescribeZoneVpcTree", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeZoneVpcTreeResponse struct {
	RequestId goaliyun.String
	Zones     DescribeZoneVpcTreeZoneList
}

type DescribeZoneVpcTreeZone struct {
	ZoneId          goaliyun.String
	ZoneName        goaliyun.String
	Remark          goaliyun.String
	RecordCount     goaliyun.Integer
	CreateTime      goaliyun.String
	CreateTimestamp goaliyun.Integer
	UpdateTime      goaliyun.String
	UpdateTimestamp goaliyun.Integer
	IsPtr           bool
	Vpcs            DescribeZoneVpcTreeVpcList
}

type DescribeZoneVpcTreeVpc struct {
	RegionId   goaliyun.String
	RegionName goaliyun.String
	VpcId      goaliyun.String
	VpcName    goaliyun.String
}

type DescribeZoneVpcTreeZoneList []DescribeZoneVpcTreeZone

func (list *DescribeZoneVpcTreeZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZoneVpcTreeZone)
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

type DescribeZoneVpcTreeVpcList []DescribeZoneVpcTreeVpc

func (list *DescribeZoneVpcTreeVpcList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZoneVpcTreeVpc)
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
