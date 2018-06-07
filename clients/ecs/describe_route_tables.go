package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRouteTablesRequest struct {
	RouterType           string `position:"Query" name:"RouterType"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RouteTableName       string `position:"Query" name:"RouteTableName"`
	VRouterId            string `position:"Query" name:"VRouterId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RouteTableId         string `position:"Query" name:"RouteTableId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRouteTablesRequest) Invoke(client goaliyun.Client) (*DescribeRouteTablesResponse, error) {
	resp := &DescribeRouteTablesResponse{}
	err := client.Request("ecs", "DescribeRouteTables", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRouteTablesResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	RouteTables DescribeRouteTablesRouteTableList
}

type DescribeRouteTablesRouteTable struct {
	VRouterId      goaliyun.String
	RouteTableId   goaliyun.String
	RouteTableType goaliyun.String
	CreationTime   goaliyun.String
	RouteEntrys    DescribeRouteTablesRouteEntryList
}

type DescribeRouteTablesRouteEntry struct {
	RouteTableId         goaliyun.String
	DestinationCidrBlock goaliyun.String
	Type                 goaliyun.String
	Status               goaliyun.String
	InstanceId           goaliyun.String
	NextHopType          goaliyun.String
	NextHops             DescribeRouteTablesNextHopList
}

type DescribeRouteTablesNextHop struct {
	NextHopType goaliyun.String
	NextHopId   goaliyun.String
	Enabled     goaliyun.Integer
	Weight      goaliyun.Integer
}

type DescribeRouteTablesRouteTableList []DescribeRouteTablesRouteTable

func (list *DescribeRouteTablesRouteTableList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTablesRouteTable)
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

type DescribeRouteTablesRouteEntryList []DescribeRouteTablesRouteEntry

func (list *DescribeRouteTablesRouteEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTablesRouteEntry)
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

type DescribeRouteTablesNextHopList []DescribeRouteTablesNextHop

func (list *DescribeRouteTablesNextHopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTablesNextHop)
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
