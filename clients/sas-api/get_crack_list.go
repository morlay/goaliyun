package sas_api

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetCrackListRequest struct {
	Start    int64  `position:"Query" name:"Start"`
	Limit    int64  `position:"Query" name:"Limit"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetCrackListRequest) Invoke(client goaliyun.Client) (*GetCrackListResponse, error) {
	resp := &GetCrackListResponse{}
	err := client.Request("sas-api", "GetCrackList", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetCrackListResponse struct {
	Code    goaliyun.String
	Message goaliyun.String
	Success bool
	Data    GetCrackListData
}

type GetCrackListData struct {
	Total goaliyun.Integer
	Items GetCrackListItemList
}

type GetCrackListItem struct {
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
	Details            GetCrackListDetailList
}

type GetCrackListDetail struct {
	Value    goaliyun.String
	Type     goaliyun.Integer
	Label    goaliyun.String
	LinkText goaliyun.String
}

type GetCrackListItemList []GetCrackListItem

func (list *GetCrackListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetCrackListItem)
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

type GetCrackListDetailList []GetCrackListDetail

func (list *GetCrackListDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetCrackListDetail)
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
