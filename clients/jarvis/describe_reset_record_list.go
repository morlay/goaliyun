package jarvis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeResetRecordListRequest struct {
	SrcIP       string `position:"Query" name:"SrcIP"`
	Period      string `position:"Query" name:"Period"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	PageSize    int64  `position:"Query" name:"PageSize"`
	CurrentPage int64  `position:"Query" name:"CurrentPage"`
	DstIP       string `position:"Query" name:"DstIP"`
	Region      string `position:"Query" name:"Region"`
	Lang        string `position:"Query" name:"Lang"`
	SourceCode  string `position:"Query" name:"SourceCode"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DescribeResetRecordListRequest) Invoke(client goaliyun.Client) (*DescribeResetRecordListResponse, error) {
	resp := &DescribeResetRecordListResponse{}
	err := client.Request("jarvis", "DescribeResetRecordList", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeResetRecordListResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
	DataList  DescribeResetRecordListDataList
	PageInfo  DescribeResetRecordListPageInfo
}

type DescribeResetRecordListData struct {
	PunishType   goaliyun.String
	DstIP        goaliyun.String
	PunishResult goaliyun.String
	DstPort      goaliyun.Integer
	SrcIP        goaliyun.String
	PunishCount  goaliyun.Integer
}

type DescribeResetRecordListPageInfo struct {
	Total       goaliyun.Integer
	PageSize    goaliyun.Integer
	CurrentPage goaliyun.Integer
}

type DescribeResetRecordListDataList []DescribeResetRecordListData

func (list *DescribeResetRecordListDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeResetRecordListData)
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
