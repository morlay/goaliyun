package jarvis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRiskTrendRequest struct {
	SourceIp      string `position:"Query" name:"SourceIp"`
	QueryProduct  string `position:"Query" name:"QueryProduct"`
	Lang          string `position:"Query" name:"Lang"`
	Peroid        string `position:"Query" name:"Peroid"`
	SourceCode    string `position:"Query" name:"SourceCode"`
	QueryRegionId string `position:"Query" name:"QueryRegionId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeRiskTrendRequest) Invoke(client goaliyun.Client) (*DescribeRiskTrendResponse, error) {
	resp := &DescribeRiskTrendResponse{}
	err := client.Request("jarvis", "DescribeRiskTrend", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRiskTrendResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.String
	Module     goaliyun.String
	DataList   DescribeRiskTrendDataItemList
}

type DescribeRiskTrendDataItem struct {
	UpdateTime     goaliyun.String
	NewRiskCount   goaliyun.Integer
	TotalRiskCount goaliyun.Integer
}

type DescribeRiskTrendDataItemList []DescribeRiskTrendDataItem

func (list *DescribeRiskTrendDataItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRiskTrendDataItem)
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
