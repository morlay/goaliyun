package sas_api

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetThreatAnalyseListRequest struct {
	Start    int64  `position:"Query" name:"Start"`
	Limit    int64  `position:"Query" name:"Limit"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetThreatAnalyseListRequest) Invoke(client goaliyun.Client) (*GetThreatAnalyseListResponse, error) {
	resp := &GetThreatAnalyseListResponse{}
	err := client.Request("sas-api", "GetThreatAnalyseList", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetThreatAnalyseListResponse struct {
	Code    goaliyun.String
	Message goaliyun.String
	Success bool
	Data    GetThreatAnalyseListData
}

type GetThreatAnalyseListData struct {
	Total goaliyun.Integer
	Items GetThreatAnalyseListItemList
}

type GetThreatAnalyseListItem struct {
	Id                 goaliyun.Integer
	Level              goaliyun.String
	Uuid               goaliyun.String
	Product            goaliyun.String
	VmIp               goaliyun.String
	Url                goaliyun.String
	Method             goaliyun.String
	SourceIp           goaliyun.String
	SourceUuid         goaliyun.String
	SourceDomain       goaliyun.String
	SourcePort         goaliyun.Integer
	SourceLocal        goaliyun.String
	DstIp              goaliyun.String
	DstUuid            goaliyun.String
	DstDomain          goaliyun.String
	DstPort            goaliyun.Integer
	DstLocal           goaliyun.String
	AttackCount        goaliyun.Integer
	ThreatRate         goaliyun.String
	AttackStartTime    goaliyun.Integer
	AttackEndTime      goaliyun.Integer
	AttackCategory     goaliyun.Integer
	AttackCategoryName goaliyun.String
	AttackType         goaliyun.String
	AttackTypeName     goaliyun.String
	AttackStatus       goaliyun.Integer
	AttackSource       goaliyun.String
	Details            GetThreatAnalyseListDetailList
}

type GetThreatAnalyseListDetail struct {
	Value    goaliyun.String
	Type     goaliyun.Integer
	Label    goaliyun.String
	LinkText goaliyun.String
}

type GetThreatAnalyseListItemList []GetThreatAnalyseListItem

func (list *GetThreatAnalyseListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThreatAnalyseListItem)
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

type GetThreatAnalyseListDetailList []GetThreatAnalyseListDetail

func (list *GetThreatAnalyseListDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThreatAnalyseListDetail)
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
