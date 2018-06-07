package jarvis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAccessWhiteListGroupRequest struct {
	SrcIP         string `position:"Query" name:"SrcIP"`
	SourceIp      string `position:"Query" name:"SourceIp"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	CurrentPage   int64  `position:"Query" name:"CurrentPage"`
	WhiteListType int64  `position:"Query" name:"WhiteListType"`
	DstIP         string `position:"Query" name:"DstIP"`
	Lang          string `position:"Query" name:"Lang"`
	Status        string `position:"Query" name:"Status"`
	SourceCode    string `position:"Query" name:"SourceCode"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeAccessWhiteListGroupRequest) Invoke(client goaliyun.Client) (*DescribeAccessWhiteListGroupResponse, error) {
	resp := &DescribeAccessWhiteListGroupResponse{}
	err := client.Request("jarvis", "DescribeAccessWhiteListGroup", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAccessWhiteListGroupResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
	DataList  DescribeAccessWhiteListGroupDataList
	PageInfo  DescribeAccessWhiteListGroupPageInfo
}

type DescribeAccessWhiteListGroupData struct {
	Status        goaliyun.String
	GmtCreate     goaliyun.String
	GmtRealExpire goaliyun.String
	SrcIP         goaliyun.String
	AutoConfig    goaliyun.Integer
	GroupId       goaliyun.Integer
	Items         DescribeAccessWhiteListGroupItemList
}

type DescribeAccessWhiteListGroupItem struct {
	IP       goaliyun.String
	RegionId goaliyun.String
}

type DescribeAccessWhiteListGroupPageInfo struct {
	Total       goaliyun.Integer
	PageSize    goaliyun.Integer
	CurrentPage goaliyun.Integer
}

type DescribeAccessWhiteListGroupDataList []DescribeAccessWhiteListGroupData

func (list *DescribeAccessWhiteListGroupDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessWhiteListGroupData)
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

type DescribeAccessWhiteListGroupItemList []DescribeAccessWhiteListGroupItem

func (list *DescribeAccessWhiteListGroupItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessWhiteListGroupItem)
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
