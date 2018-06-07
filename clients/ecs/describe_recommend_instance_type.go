package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRecommendInstanceTypeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Channel              string `position:"Query" name:"Channel"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Operator             string `position:"Query" name:"Operator"`
	Token                string `position:"Query" name:"Token"`
	Scene                string `position:"Query" name:"Scene"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRecommendInstanceTypeRequest) Invoke(client goaliyun.Client) (*DescribeRecommendInstanceTypeResponse, error) {
	resp := &DescribeRecommendInstanceTypeResponse{}
	err := client.Request("ecs", "DescribeRecommendInstanceType", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRecommendInstanceTypeResponse struct {
	RequestId goaliyun.String
	Data      DescribeRecommendInstanceTypeRecommendInstanceTypeList
}

type DescribeRecommendInstanceTypeRecommendInstanceType struct {
	RegionNo      goaliyun.String
	CommodityCode goaliyun.String
	Scene         goaliyun.String
	Zones         DescribeRecommendInstanceTypeZoneList
	InstanceType  DescribeRecommendInstanceTypeInstanceType
}

type DescribeRecommendInstanceTypeZone struct {
	ZoneNo       goaliyun.String
	NetworkTypes DescribeRecommendInstanceTypeNetworkTypeList
}

type DescribeRecommendInstanceTypeInstanceType struct {
	Generation         goaliyun.String
	InstanceTypeFamily goaliyun.String
	InstanceType       goaliyun.String
	SupportIoOptimized goaliyun.String
	Cores              goaliyun.Integer
	Memory             goaliyun.Integer
}

type DescribeRecommendInstanceTypeRecommendInstanceTypeList []DescribeRecommendInstanceTypeRecommendInstanceType

func (list *DescribeRecommendInstanceTypeRecommendInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecommendInstanceTypeRecommendInstanceType)
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

type DescribeRecommendInstanceTypeZoneList []DescribeRecommendInstanceTypeZone

func (list *DescribeRecommendInstanceTypeZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecommendInstanceTypeZone)
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

type DescribeRecommendInstanceTypeNetworkTypeList []goaliyun.String

func (list *DescribeRecommendInstanceTypeNetworkTypeList) UnmarshalJSON(data []byte) error {
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
