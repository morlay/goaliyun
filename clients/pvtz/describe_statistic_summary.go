package pvtz

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeStatisticSummaryRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeStatisticSummaryRequest) Invoke(client goaliyun.Client) (*DescribeStatisticSummaryResponse, error) {
	resp := &DescribeStatisticSummaryResponse{}
	err := client.Request("pvtz", "DescribeStatisticSummary", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeStatisticSummaryResponse struct {
	RequestId       goaliyun.String
	TotalCount      goaliyun.Integer
	ZoneRequestTops DescribeStatisticSummaryZoneRequestTopList
	VpcRequestTops  DescribeStatisticSummaryVpcRequestTopList
}

type DescribeStatisticSummaryZoneRequestTop struct {
	ZoneName     goaliyun.String
	RequestCount goaliyun.Integer
}

type DescribeStatisticSummaryVpcRequestTop struct {
	RegionId     goaliyun.String
	VpcId        goaliyun.String
	TunnelId     goaliyun.String
	RequestCount goaliyun.Integer
	RegionName   goaliyun.String
}

type DescribeStatisticSummaryZoneRequestTopList []DescribeStatisticSummaryZoneRequestTop

func (list *DescribeStatisticSummaryZoneRequestTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStatisticSummaryZoneRequestTop)
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

type DescribeStatisticSummaryVpcRequestTopList []DescribeStatisticSummaryVpcRequestTop

func (list *DescribeStatisticSummaryVpcRequestTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStatisticSummaryVpcRequestTop)
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
