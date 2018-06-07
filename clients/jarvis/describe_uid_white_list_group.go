package jarvis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeUidWhiteListGroupRequest struct {
	SourceIp      string `position:"Query" name:"SourceIp"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	CurrentPage   int64  `position:"Query" name:"CurrentPage"`
	WhiteListType int64  `position:"Query" name:"WhiteListType"`
	DstIP         string `position:"Query" name:"DstIP"`
	Lang          string `position:"Query" name:"Lang"`
	SrcUid        string `position:"Query" name:"SrcUid"`
	Status        string `position:"Query" name:"Status"`
	SourceCode    string `position:"Query" name:"SourceCode"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeUidWhiteListGroupRequest) Invoke(client goaliyun.Client) (*DescribeUidWhiteListGroupResponse, error) {
	resp := &DescribeUidWhiteListGroupResponse{}
	err := client.Request("jarvis", "DescribeUidWhiteListGroup", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUidWhiteListGroupResponse struct {
	RequestId   goaliyun.String
	Module      goaliyun.String
	DataList    DescribeUidWhiteListGroupDataList
	ProductList DescribeUidWhiteListGroupProductListList
	PageInfo    DescribeUidWhiteListGroupPageInfo
}

type DescribeUidWhiteListGroupData struct {
	Status        goaliyun.String
	GmtCreate     goaliyun.String
	GmtRealExpire goaliyun.String
	SrcUid        goaliyun.String
	AutoConfig    goaliyun.Integer
	GroupId       goaliyun.Integer
	Items         DescribeUidWhiteListGroupItemList
}

type DescribeUidWhiteListGroupItem struct {
	IP       goaliyun.String
	RegionId goaliyun.String
}

type DescribeUidWhiteListGroupPageInfo struct {
	Total       goaliyun.Integer
	PageSize    goaliyun.Integer
	CurrentPage goaliyun.Integer
}

type DescribeUidWhiteListGroupDataList []DescribeUidWhiteListGroupData

func (list *DescribeUidWhiteListGroupDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeUidWhiteListGroupData)
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

type DescribeUidWhiteListGroupProductListList []goaliyun.String

func (list *DescribeUidWhiteListGroupProductListList) UnmarshalJSON(data []byte) error {
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

type DescribeUidWhiteListGroupItemList []DescribeUidWhiteListGroupItem

func (list *DescribeUidWhiteListGroupItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeUidWhiteListGroupItem)
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
