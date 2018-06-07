package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindApproveServiceListRequest struct {
	ApproveLevel   string `position:"Query" name:"ApproveLevel"`
	ProjectName    string `position:"Query" name:"ProjectName"`
	ShowDelService string `position:"Query" name:"ShowDelService"`
	CsbId          int64  `position:"Query" name:"CsbId"`
	Alias          string `position:"Query" name:"Alias"`
	ServiceName    string `position:"Query" name:"ServiceName"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *FindApproveServiceListRequest) Invoke(client goaliyun.Client) (*FindApproveServiceListResponse, error) {
	resp := &FindApproveServiceListResponse{}
	err := client.Request("csb", "FindApproveServiceList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindApproveServiceListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      FindApproveServiceListData
}

type FindApproveServiceListData struct {
	Total       goaliyun.Integer
	PageNumber  goaliyun.Integer
	CurrentPage goaliyun.Integer
	ServiceList FindApproveServiceListServiceList
}

type FindApproveServiceListService struct {
	AllVisiable    bool
	CasTargets     goaliyun.String
	CreateTime     goaliyun.Integer
	CsbId          goaliyun.Integer
	Id             goaliyun.Integer
	InterfaceName  goaliyun.String
	ModifiedTime   goaliyun.Integer
	OwnerId        goaliyun.String
	PrincipalName  goaliyun.String
	ProjectId      goaliyun.Integer
	ProjectName    goaliyun.String
	Qps            goaliyun.Integer
	Scope          goaliyun.String
	ServiceName    goaliyun.String
	ServiceVersion goaliyun.String
	SkipAuth       bool
	StatisticName  goaliyun.String
	Status         goaliyun.Integer
	UserId         goaliyun.String
}

type FindApproveServiceListServiceList []FindApproveServiceListService

func (list *FindApproveServiceListServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindApproveServiceListService)
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
