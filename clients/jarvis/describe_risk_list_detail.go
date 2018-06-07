package jarvis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRiskListDetailRequest struct {
	RiskType      string `position:"Query" name:"RiskType"`
	SourceIp      string `position:"Query" name:"SourceIp"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	QueryProduct  string `position:"Query" name:"QueryProduct"`
	CurrentPage   int64  `position:"Query" name:"CurrentPage"`
	RiskDescribe  string `position:"Query" name:"RiskDescribe"`
	Lang          string `position:"Query" name:"Lang"`
	SrcUid        int64  `position:"Query" name:"SrcUid"`
	SourceCode    string `position:"Query" name:"SourceCode"`
	QueryRegionId string `position:"Query" name:"QueryRegionId"`
	Status        string `position:"Query" name:"Status"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeRiskListDetailRequest) Invoke(client goaliyun.Client) (*DescribeRiskListDetailResponse, error) {
	resp := &DescribeRiskListDetailResponse{}
	err := client.Request("jarvis", "DescribeRiskListDetail", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRiskListDetailResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
	DataList  DescribeRiskListDetailDataList
	PageInfo  DescribeRiskListDetailPageInfo
}

type DescribeRiskListDetailData struct {
	RiskId           goaliyun.Integer
	UpdateTime       goaliyun.String
	RiskDescribe     goaliyun.String
	RiskType         goaliyun.String
	RiskInstance     goaliyun.String
	Product          goaliyun.String
	RegionId         goaliyun.String
	TacticsName      goaliyun.String
	Status           goaliyun.String
	IgnoreTime       goaliyun.String
	EcsSecGroupRisk  DescribeRiskListDetailEcsSecGroupRiskItemList
	RdsWhitelistRisk DescribeRiskListDetailRdsWhitelistRiskItemList
}

type DescribeRiskListDetailEcsSecGroupRiskItem struct {
	NetType      goaliyun.String
	Direction    goaliyun.String
	DstPortRange goaliyun.String
	SrcIpRange   goaliyun.String
}

type DescribeRiskListDetailRdsWhitelistRiskItem struct {
	RdsWhitelistGroup goaliyun.String
	RiskDetail        goaliyun.String
}

type DescribeRiskListDetailPageInfo struct {
	Total       goaliyun.Integer
	PageSize    goaliyun.Integer
	CurrentPage goaliyun.Integer
}

type DescribeRiskListDetailDataList []DescribeRiskListDetailData

func (list *DescribeRiskListDetailDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRiskListDetailData)
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

type DescribeRiskListDetailEcsSecGroupRiskItemList []DescribeRiskListDetailEcsSecGroupRiskItem

func (list *DescribeRiskListDetailEcsSecGroupRiskItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRiskListDetailEcsSecGroupRiskItem)
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

type DescribeRiskListDetailRdsWhitelistRiskItemList []DescribeRiskListDetailRdsWhitelistRiskItem

func (list *DescribeRiskListDetailRdsWhitelistRiskItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRiskListDetailRdsWhitelistRiskItem)
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
