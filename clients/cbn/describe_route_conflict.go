package cbn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRouteConflictRequest struct {
	ChildInstanceId           string `position:"Query" name:"ChildInstanceId"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	DestinationCidrBlock      string `position:"Query" name:"DestinationCidrBlock"`
	PageSize                  int64  `position:"Query" name:"PageSize"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	ChildInstanceType         string `position:"Query" name:"ChildInstanceType"`
	ChildInstanceRouteTableId string `position:"Query" name:"ChildInstanceRouteTableId"`
	PageNumber                int64  `position:"Query" name:"PageNumber"`
	ChildInstanceRegionId     string `position:"Query" name:"ChildInstanceRegionId"`
	RegionId                  string `position:"Query" name:"RegionId"`
}

func (req *DescribeRouteConflictRequest) Invoke(client goaliyun.Client) (*DescribeRouteConflictResponse, error) {
	resp := &DescribeRouteConflictResponse{}
	err := client.Request("cbn", "DescribeRouteConflict", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRouteConflictResponse struct {
	RequestId      goaliyun.String
	PageNumber     goaliyun.Integer
	TotalCount     goaliyun.Integer
	PageSize       goaliyun.Integer
	RouteConflicts DescribeRouteConflictRouteConflictList
}

type DescribeRouteConflictRouteConflict struct {
	DestinationCidrBlock goaliyun.String
	RegionId             goaliyun.String
	InstanceId           goaliyun.String
	InstanceType         goaliyun.String
	Status               goaliyun.String
}

type DescribeRouteConflictRouteConflictList []DescribeRouteConflictRouteConflict

func (list *DescribeRouteConflictRouteConflictList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteConflictRouteConflict)
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
