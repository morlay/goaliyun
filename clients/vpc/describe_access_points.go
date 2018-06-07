package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAccessPointsRequest struct {
	Filters              *DescribeAccessPointsFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                           `position:"Query" name:"ResourceOwnerId"`
	HostOperator         string                          `position:"Query" name:"HostOperator"`
	ResourceOwnerAccount string                          `position:"Query" name:"ResourceOwnerAccount"`
	Name                 string                          `position:"Query" name:"Name"`
	PageSize             int64                           `position:"Query" name:"PageSize"`
	OwnerId              int64                           `position:"Query" name:"OwnerId"`
	Type                 string                          `position:"Query" name:"Type"`
	PageNumber           int64                           `position:"Query" name:"PageNumber"`
	RegionId             string                          `position:"Query" name:"RegionId"`
}

func (req *DescribeAccessPointsRequest) Invoke(client goaliyun.Client) (*DescribeAccessPointsResponse, error) {
	resp := &DescribeAccessPointsResponse{}
	err := client.Request("vpc", "DescribeAccessPoints", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAccessPointsFilter struct {
	Key    string                         `name:"Key"`
	Values *DescribeAccessPointsValueList `type:"Repeated" name:"Value"`
}

type DescribeAccessPointsResponse struct {
	RequestId      goaliyun.String
	PageNumber     goaliyun.Integer
	PageSize       goaliyun.Integer
	TotalCount     goaliyun.Integer
	AccessPointSet DescribeAccessPointsAccessPointTypeList
}

type DescribeAccessPointsAccessPointType struct {
	AccessPointId    goaliyun.String
	Status           goaliyun.String
	Type             goaliyun.String
	AttachedRegionNo goaliyun.String
	Location         goaliyun.String
	HostOperator     goaliyun.String
	Name             goaliyun.String
	Description      goaliyun.String
}

type DescribeAccessPointsFilterList []DescribeAccessPointsFilter

func (list *DescribeAccessPointsFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessPointsFilter)
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

type DescribeAccessPointsValueList []string

func (list *DescribeAccessPointsValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeAccessPointsAccessPointTypeList []DescribeAccessPointsAccessPointType

func (list *DescribeAccessPointsAccessPointTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessPointsAccessPointType)
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
