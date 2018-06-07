package cbn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeGeographicRegionMembershipRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	GeographicRegionId   string `position:"Query" name:"GeographicRegionId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeGeographicRegionMembershipRequest) Invoke(client goaliyun.Client) (*DescribeGeographicRegionMembershipResponse, error) {
	resp := &DescribeGeographicRegionMembershipResponse{}
	err := client.Request("cbn", "DescribeGeographicRegionMembership", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeGeographicRegionMembershipResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	RegionIds  DescribeGeographicRegionMembershipRegionIdList
}

type DescribeGeographicRegionMembershipRegionId struct {
	RegionId goaliyun.String
}

type DescribeGeographicRegionMembershipRegionIdList []DescribeGeographicRegionMembershipRegionId

func (list *DescribeGeographicRegionMembershipRegionIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGeographicRegionMembershipRegionId)
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
