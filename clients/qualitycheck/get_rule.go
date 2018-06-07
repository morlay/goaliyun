package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetRuleRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetRuleRequest) Invoke(client goaliyun.Client) (*GetRuleResponse, error) {
	resp := &GetRuleResponse{}
	err := client.Request("qualitycheck", "GetRule", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetRuleResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      GetRuleData
}

type GetRuleData struct {
	Rules GetRuleRuleInfoList
}

type GetRuleRuleInfo struct {
	Rid                      goaliyun.String
	RuleLambda               goaliyun.String
	Name                     goaliyun.String
	Type                     goaliyun.Integer
	Status                   goaliyun.Integer
	IsDelete                 goaliyun.Integer
	StartTime                goaliyun.String
	EndTime                  goaliyun.String
	Weight                   goaliyun.String
	IsOnline                 goaliyun.Integer
	CreateEmpid              goaliyun.String
	CreateTime               goaliyun.String
	LastUpdateTime           goaliyun.String
	LastUpdateEmpid          goaliyun.String
	Comments                 goaliyun.String
	AutoReview               goaliyun.Integer
	RuleScoreType            goaliyun.Integer
	ScoreName                goaliyun.String
	ScoreSubName             goaliyun.String
	BusinessCategoryNameList GetRuleBusinessCategoryNameListList
}

type GetRuleRuleInfoList []GetRuleRuleInfo

func (list *GetRuleRuleInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleRuleInfo)
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

type GetRuleBusinessCategoryNameListList []goaliyun.String

func (list *GetRuleBusinessCategoryNameListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
