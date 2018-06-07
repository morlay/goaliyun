package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRouteTableListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RouterType           string `position:"Query" name:"RouterType"`
	RouteTableName       string `position:"Query" name:"RouteTableName"`
	RouterId             string `position:"Query" name:"RouterId"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	RouteTableId         string `position:"Query" name:"RouteTableId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRouteTableListRequest) Invoke(client goaliyun.Client) (*DescribeRouteTableListResponse, error) {
	resp := &DescribeRouteTableListResponse{}
	err := client.Request("vpc", "DescribeRouteTableList", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRouteTableListResponse struct {
	RequestId       goaliyun.String
	Code            goaliyun.String
	Message         goaliyun.String
	Success         bool
	PageSize        goaliyun.Integer
	PageNumber      goaliyun.Integer
	TotalCount      goaliyun.Integer
	RouterTableList DescribeRouteTableListRouterTableListTypeList
}

type DescribeRouteTableListRouterTableListType struct {
	VpcId          goaliyun.String
	RouterType     goaliyun.String
	RouterId       goaliyun.String
	RouteTableId   goaliyun.String
	RouteTableName goaliyun.String
	RouteTableType goaliyun.String
	Description    goaliyun.String
	CreationTime   goaliyun.String
}

type DescribeRouteTableListRouterTableListTypeList []DescribeRouteTableListRouterTableListType

func (list *DescribeRouteTableListRouterTableListTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTableListRouterTableListType)
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
