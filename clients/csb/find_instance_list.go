package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindInstanceListRequest struct {
	SearchTxt string `position:"Query" name:"SearchTxt"`
	CsbId     int64  `position:"Query" name:"CsbId"`
	PageNum   int64  `position:"Query" name:"PageNum"`
	Status    int64  `position:"Query" name:"Status"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *FindInstanceListRequest) Invoke(client goaliyun.Client) (*FindInstanceListResponse, error) {
	resp := &FindInstanceListResponse{}
	err := client.Request("csb", "FindInstanceList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindInstanceListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      FindInstanceListData
}

type FindInstanceListData struct {
	CurrentPage goaliyun.Integer
	PageNumber  goaliyun.Integer
	ItemList    FindInstanceListItemList
}

type FindInstanceListItem struct {
	Description      goaliyun.String
	FrontStatus      goaliyun.String
	GmtCreate        goaliyun.Integer
	GmtModified      goaliyun.Integer
	Id               goaliyun.Integer
	InstanceCategory goaliyun.Integer
	Name             goaliyun.String
	StatusCode       goaliyun.Integer
	Visible          bool
	VpcName          goaliyun.String
}

type FindInstanceListItemList []FindInstanceListItem

func (list *FindInstanceListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindInstanceListItem)
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
