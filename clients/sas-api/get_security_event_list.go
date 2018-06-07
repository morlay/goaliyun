package sas_api

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetSecurityEventListRequest struct {
	Start    int64  `position:"Query" name:"Start"`
	Limit    int64  `position:"Query" name:"Limit"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetSecurityEventListRequest) Invoke(client goaliyun.Client) (*GetSecurityEventListResponse, error) {
	resp := &GetSecurityEventListResponse{}
	err := client.Request("sas-api", "GetSecurityEventList", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetSecurityEventListResponse struct {
	Code    goaliyun.String
	Message goaliyun.String
	Success bool
	Data    GetSecurityEventListData
}

type GetSecurityEventListData struct {
	Total goaliyun.Integer
	Items GetSecurityEventListItemList
}

type GetSecurityEventListItem struct {
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
	Details            GetSecurityEventListDetailList
}

type GetSecurityEventListDetail struct {
	Value    goaliyun.String
	Type     goaliyun.Integer
	Label    goaliyun.String
	LinkText goaliyun.String
}

type GetSecurityEventListItemList []GetSecurityEventListItem

func (list *GetSecurityEventListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSecurityEventListItem)
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

type GetSecurityEventListDetailList []GetSecurityEventListDetail

func (list *GetSecurityEventListDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSecurityEventListDetail)
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
