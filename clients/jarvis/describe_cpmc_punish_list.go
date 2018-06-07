package jarvis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCpmcPunishListRequest struct {
	SrcIP        string `position:"Query" name:"SrcIP"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	CurrentPage  int64  `position:"Query" name:"CurrentPage"`
	PunishStatus string `position:"Query" name:"PunishStatus"`
	Lang         string `position:"Query" name:"Lang"`
	SourceCode   string `position:"Query" name:"SourceCode"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeCpmcPunishListRequest) Invoke(client goaliyun.Client) (*DescribeCpmcPunishListResponse, error) {
	resp := &DescribeCpmcPunishListResponse{}
	err := client.Request("jarvis", "DescribeCpmcPunishList", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCpmcPunishListResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
	DataList  DescribeCpmcPunishListDataList
	PageInfo  DescribeCpmcPunishListPageInfo
}

type DescribeCpmcPunishListData struct {
	GmtCreate    goaliyun.String
	SrcPort      goaliyun.Integer
	FeedBack     goaliyun.Integer
	GmtExpire    goaliyun.String
	PunishType   goaliyun.String
	DstIP        goaliyun.String
	PunishResult goaliyun.String
	RegionId     goaliyun.String
	DstPort      goaliyun.Integer
	Protocol     goaliyun.String
	SrcIP        goaliyun.String
	Reason       goaliyun.String
}

type DescribeCpmcPunishListPageInfo struct {
	Total       goaliyun.Integer
	PageSize    goaliyun.Integer
	CurrentPage goaliyun.Integer
}

type DescribeCpmcPunishListDataList []DescribeCpmcPunishListData

func (list *DescribeCpmcPunishListDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCpmcPunishListData)
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
