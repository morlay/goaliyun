package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindOrderedListRequest struct {
	ProjectName  string `position:"Query" name:"ProjectName"`
	ShowDelOrder string `position:"Query" name:"ShowDelOrder"`
	CsbId        int64  `position:"Query" name:"CsbId"`
	Alias        string `position:"Query" name:"Alias"`
	ServiceName  string `position:"Query" name:"ServiceName"`
	PageNum      int64  `position:"Query" name:"PageNum"`
	ServiceId    int64  `position:"Query" name:"ServiceId"`
	Status       string `position:"Query" name:"Status"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *FindOrderedListRequest) Invoke(client goaliyun.Client) (*FindOrderedListResponse, error) {
	resp := &FindOrderedListResponse{}
	err := client.Request("csb", "FindOrderedList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindOrderedListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      FindOrderedListData
}

type FindOrderedListData struct {
	CurrentPage goaliyun.Integer
	PageNumber  goaliyun.Integer
	OrderList   FindOrderedListOrderList
}

type FindOrderedListOrder struct {
	Alias                 goaliyun.String
	ProjectName           goaliyun.String
	ServiceName           goaliyun.String
	ServiceVersion        goaliyun.String
	OrderStatus           goaliyun.Integer
	AliveOrderCount       goaliyun.Integer
	GmtCreate             goaliyun.Integer
	MaxRT                 goaliyun.Integer
	MinRT                 goaliyun.Integer
	ServiceId             goaliyun.String
	ServiceStatus         goaliyun.Integer
	ErrorTypeCatagoryList FindOrderedListErrorTypeCatagoryList
	Orders                FindOrderedListOrder1List
	Total                 FindOrderedListTotal
}

type FindOrderedListErrorTypeCatagory struct {
	Name     goaliyun.String
	Total    goaliyun.Integer
	ErrorNum goaliyun.Integer
}

type FindOrderedListOrder1 struct {
	Alias           goaliyun.String
	ApproveComments goaliyun.String
	CsbId           goaliyun.Integer
	GmtCreate       goaliyun.Integer
	GmtModified     goaliyun.Integer
	GroupName       goaliyun.String
	Id              goaliyun.Integer
	ProjectName     goaliyun.String
	ServiceId       goaliyun.Integer
	ServiceName     goaliyun.String
	ServiceStatus   goaliyun.Integer
	ServiceVersion  goaliyun.String
	StatisticName   goaliyun.String
	Status          goaliyun.Integer
	UserId          goaliyun.String
	SlaInfo         FindOrderedListSlaInfo
	Total2          FindOrderedListTotal2
}

type FindOrderedListSlaInfo struct {
	Qph goaliyun.String
	Qps goaliyun.String
}

type FindOrderedListTotal2 struct {
	ErrorNum goaliyun.Integer
	Total    goaliyun.Integer
}

type FindOrderedListTotal struct {
	ErrorNum goaliyun.Integer
	Total    goaliyun.Integer
}

type FindOrderedListOrderList []FindOrderedListOrder

func (list *FindOrderedListOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindOrderedListOrder)
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

type FindOrderedListErrorTypeCatagoryList []FindOrderedListErrorTypeCatagory

func (list *FindOrderedListErrorTypeCatagoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindOrderedListErrorTypeCatagory)
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

type FindOrderedListOrder1List []FindOrderedListOrder1

func (list *FindOrderedListOrder1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindOrderedListOrder1)
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
