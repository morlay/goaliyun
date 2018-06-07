package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetRuleDetailRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetRuleDetailRequest) Invoke(client goaliyun.Client) (*GetRuleDetailResponse, error) {
	resp := &GetRuleDetailResponse{}
	err := client.Request("qualitycheck", "GetRuleDetail", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetRuleDetailResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      GetRuleDetailData
}

type GetRuleDetailData struct {
	Conditions GetRuleDetailConditionBasicInfoList
	Rules      GetRuleDetailRuleBasicInfoList
}

type GetRuleDetailConditionBasicInfo struct {
	ConditionInfoCid goaliyun.String
	OperLambda       goaliyun.String
	Operators        GetRuleDetailOperatorBasicInfoList
	CheckRange       GetRuleDetailCheckRange
}

type GetRuleDetailOperatorBasicInfo struct {
	Oid      goaliyun.String
	Type     goaliyun.String
	OperName goaliyun.String
	Param    GetRuleDetailParam
}

type GetRuleDetailParam struct {
	Regex         goaliyun.String
	Phrase        goaliyun.String
	Interval      goaliyun.Integer
	Threshold     goaliyun.Float
	InSentence    bool
	Target        goaliyun.Integer
	FromEnd       bool
	DifferentRole bool
	TargetRole    goaliyun.String
	OperKeyWords  GetRuleDetailOperKeyWordList
	References    GetRuleDetailReferenceList
}

type GetRuleDetailCheckRange struct {
	Role   goaliyun.String
	Anchor GetRuleDetailAnchor
	Range  GetRuleDetailRange
}

type GetRuleDetailAnchor struct {
	AnchorCid goaliyun.String
	Location  goaliyun.String
	HitTime   goaliyun.Integer
}

type GetRuleDetailRange struct {
	From goaliyun.Integer
	To   goaliyun.Integer
}

type GetRuleDetailRuleBasicInfo struct {
	Rid                goaliyun.String
	RuleLambda         goaliyun.String
	BusinessCategories GetRuleDetailBusinessCategoryBasicInfoList
	Triggers           GetRuleDetailTriggerList
}

type GetRuleDetailBusinessCategoryBasicInfo struct {
	Bid          goaliyun.Integer
	ServiceType  goaliyun.Integer
	BusinessName goaliyun.String
}

type GetRuleDetailConditionBasicInfoList []GetRuleDetailConditionBasicInfo

func (list *GetRuleDetailConditionBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleDetailConditionBasicInfo)
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

type GetRuleDetailRuleBasicInfoList []GetRuleDetailRuleBasicInfo

func (list *GetRuleDetailRuleBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleDetailRuleBasicInfo)
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

type GetRuleDetailOperatorBasicInfoList []GetRuleDetailOperatorBasicInfo

func (list *GetRuleDetailOperatorBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleDetailOperatorBasicInfo)
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

type GetRuleDetailOperKeyWordList []goaliyun.String

func (list *GetRuleDetailOperKeyWordList) UnmarshalJSON(data []byte) error {
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

type GetRuleDetailReferenceList []goaliyun.String

func (list *GetRuleDetailReferenceList) UnmarshalJSON(data []byte) error {
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

type GetRuleDetailBusinessCategoryBasicInfoList []GetRuleDetailBusinessCategoryBasicInfo

func (list *GetRuleDetailBusinessCategoryBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleDetailBusinessCategoryBasicInfo)
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

type GetRuleDetailTriggerList []goaliyun.String

func (list *GetRuleDetailTriggerList) UnmarshalJSON(data []byte) error {
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
