package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVRoutersRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VRouterId            string `position:"Query" name:"VRouterId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVRoutersRequest) Invoke(client goaliyun.Client) (*DescribeVRoutersResponse, error) {
	resp := &DescribeVRoutersResponse{}
	err := client.Request("ecs", "DescribeVRouters", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVRoutersResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	VRouters   DescribeVRoutersVRouterList
}

type DescribeVRoutersVRouter struct {
	RegionId      goaliyun.String
	VpcId         goaliyun.String
	VRouterName   goaliyun.String
	Description   goaliyun.String
	VRouterId     goaliyun.String
	CreationTime  goaliyun.String
	RouteTableIds DescribeVRoutersRouteTableIdList
}

type DescribeVRoutersVRouterList []DescribeVRoutersVRouter

func (list *DescribeVRoutersVRouterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVRoutersVRouter)
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

type DescribeVRoutersRouteTableIdList []goaliyun.String

func (list *DescribeVRoutersRouteTableIdList) UnmarshalJSON(data []byte) error {
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
