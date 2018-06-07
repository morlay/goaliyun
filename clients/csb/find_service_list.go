package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindServiceListRequest struct {
	ProjectName    string `position:"Query" name:"ProjectName"`
	CasShowType    int64  `position:"Query" name:"CasShowType"`
	ShowDelService string `position:"Query" name:"ShowDelService"`
	CsbId          int64  `position:"Query" name:"CsbId"`
	Alias          string `position:"Query" name:"Alias"`
	ServiceName    string `position:"Query" name:"ServiceName"`
	PageNum        int64  `position:"Query" name:"PageNum"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *FindServiceListRequest) Invoke(client goaliyun.Client) (*FindServiceListResponse, error) {
	resp := &FindServiceListResponse{}
	err := client.Request("csb", "FindServiceList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindServiceListResponse struct {
	Message   goaliyun.String
	Code      goaliyun.Integer
	RequestId goaliyun.String
	Data      FindServiceListData
}

type FindServiceListData struct {
	CurrentPage goaliyun.Integer
	PageNumber  goaliyun.Integer
	Total       goaliyun.Integer
	ServiceList FindServiceListServiceList
}

type FindServiceListService struct {
	Alias          goaliyun.String
	AllVisiable    bool
	CreateTime     goaliyun.Integer
	CsbId          goaliyun.Integer
	Description    goaliyun.String
	Id             goaliyun.Integer
	InterfaceName  goaliyun.String
	ModifiedTime   goaliyun.Integer
	OrderInfo      goaliyun.String
	OwnerId        goaliyun.String
	PrincipalName  goaliyun.String
	ProjectId      goaliyun.Integer
	ProjectName    goaliyun.String
	Scope          goaliyun.String
	ServiceName    goaliyun.String
	ServiceVersion goaliyun.String
	SkipAuth       bool
	StatisticName  goaliyun.String
	Status         goaliyun.Integer
	UserId         goaliyun.String
	CasTargets     goaliyun.String
}

type FindServiceListServiceList []FindServiceListService

func (list *FindServiceListServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindServiceListService)
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
