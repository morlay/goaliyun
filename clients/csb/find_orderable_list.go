package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindOrderableListRequest struct {
	ProjectName string `position:"Query" name:"ProjectName"`
	CsbId       int64  `position:"Query" name:"CsbId"`
	Alias       string `position:"Query" name:"Alias"`
	ServiceName string `position:"Query" name:"ServiceName"`
	PageNum     int64  `position:"Query" name:"PageNum"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *FindOrderableListRequest) Invoke(client goaliyun.Client) (*FindOrderableListResponse, error) {
	resp := &FindOrderableListResponse{}
	err := client.Request("csb", "FindOrderableList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindOrderableListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      FindOrderableListData
}

type FindOrderableListData struct {
	CurrentPage goaliyun.Integer
	PageNumber  goaliyun.Integer
	ServiceList FindOrderableListServiceList
}

type FindOrderableListService struct {
	Alias          goaliyun.String
	AllVisiable    bool
	ApproveUserId  goaliyun.String
	CasTargets     goaliyun.String
	CreateTime     goaliyun.Integer
	CsbId          goaliyun.Integer
	Id             goaliyun.Integer
	InterfaceName  goaliyun.String
	ModifiedTime   goaliyun.Integer
	OwnerId        goaliyun.String
	PrincipalName  goaliyun.String
	ProjectId      goaliyun.String
	ProjectName    goaliyun.String
	Scope          goaliyun.String
	ServiceName    goaliyun.String
	ServiceVersion goaliyun.String
	SkipAuth       bool
	StatisticName  goaliyun.String
	Status         goaliyun.Integer
	UserId         goaliyun.String
}

type FindOrderableListServiceList []FindOrderableListService

func (list *FindOrderableListServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindOrderableListService)
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
