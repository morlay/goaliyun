package jarvis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePunishListRequest struct {
	SrcIP        string `position:"Query" name:"SrcIP"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	CurrentPage  int64  `position:"Query" name:"CurrentPage"`
	PunishStatus string `position:"Query" name:"PunishStatus"`
	Lang         string `position:"Query" name:"Lang"`
	SrcUid       int64  `position:"Query" name:"SrcUid"`
	SourceCode   string `position:"Query" name:"SourceCode"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribePunishListRequest) Invoke(client goaliyun.Client) (*DescribePunishListResponse, error) {
	resp := &DescribePunishListResponse{}
	err := client.Request("jarvis", "DescribePunishList", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePunishListResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
	DataList  DescribePunishListDataList
	PageInfo  DescribePunishListPageInfo
}

type DescribePunishListData struct {
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

type DescribePunishListPageInfo struct {
	Total       goaliyun.Integer
	PageSize    goaliyun.Integer
	CurrentPage goaliyun.Integer
}

type DescribePunishListDataList []DescribePunishListData

func (list *DescribePunishListDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePunishListData)
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
