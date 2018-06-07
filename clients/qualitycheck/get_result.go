package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetResultRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetResultRequest) Invoke(client goaliyun.Client) (*GetResultResponse, error) {
	resp := &GetResultResponse{}
	err := client.Request("qualitycheck", "GetResult", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetResultResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Count     goaliyun.Integer
	Data      GetResultResultInfoList
}

type GetResultResultInfo struct {
	Tid             goaliyun.String
	AsrMsg          goaliyun.String
	Score           goaliyun.Integer
	ReviewStatus    goaliyun.Integer
	HitId           goaliyun.String
	Rules           GetResultRuleHitInfoList
	HandScoreIdList GetResultHandScoreIdListList
}

type GetResultRuleHitInfo struct {
	HitStatus     goaliyun.Integer
	Rid           goaliyun.String
	Hit           GetResultConditionHitInfoList
	ConditionInfo GetResultConditionBasicInfoList
}

type GetResultConditionHitInfo struct {
	HitKeyWords GetResultHitKeyWordList
	HitCids     GetResultHitCidList
	Phrase      GetResultPhrase
}

type GetResultHitKeyWord struct {
	Val  goaliyun.String
	Pid  goaliyun.Integer
	From goaliyun.Integer
	To   goaliyun.Integer
	Tid  goaliyun.String
}

type GetResultPhrase struct {
	Role       goaliyun.String
	Identity   goaliyun.String
	Words      goaliyun.String
	Begin      goaliyun.Integer
	End        goaliyun.Integer
	BeginTime  goaliyun.String
	HourMinSec goaliyun.String
}

type GetResultConditionBasicInfo struct {
	ConditionInfoCid goaliyun.String
	Lambda           goaliyun.String
	Operators        GetResultOperatorBasicInfoList
	CheckRange       GetResultCheckRange
}

type GetResultOperatorBasicInfo struct {
	Oid   goaliyun.String
	Type  goaliyun.String
	Name  goaliyun.String
	Param GetResultParam
}

type GetResultParam struct {
	Regex         goaliyun.String
	Phrase        goaliyun.String
	Interval      goaliyun.Integer
	Threshold     goaliyun.Float
	InSentence    bool
	Target        goaliyun.Integer
	FromEnd       bool
	DifferentRole bool
	TargetRole    goaliyun.String
	OperKeyWords  GetResultOperKeyWordList
	References    GetResultReferenceList
}

type GetResultCheckRange struct {
	Role   goaliyun.String
	Anchor GetResultAnchor
	Range  GetResultRange
}

type GetResultAnchor struct {
	AnchorCid goaliyun.String
	Location  goaliyun.String
	HitTime   goaliyun.Integer
}

type GetResultRange struct {
	From goaliyun.Integer
	To   goaliyun.Integer
}

type GetResultResultInfoList []GetResultResultInfo

func (list *GetResultResultInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultResultInfo)
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

type GetResultRuleHitInfoList []GetResultRuleHitInfo

func (list *GetResultRuleHitInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultRuleHitInfo)
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

type GetResultHandScoreIdListList []goaliyun.String

func (list *GetResultHandScoreIdListList) UnmarshalJSON(data []byte) error {
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

type GetResultConditionHitInfoList []GetResultConditionHitInfo

func (list *GetResultConditionHitInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultConditionHitInfo)
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

type GetResultConditionBasicInfoList []GetResultConditionBasicInfo

func (list *GetResultConditionBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultConditionBasicInfo)
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

type GetResultHitKeyWordList []GetResultHitKeyWord

func (list *GetResultHitKeyWordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultHitKeyWord)
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

type GetResultHitCidList []goaliyun.String

func (list *GetResultHitCidList) UnmarshalJSON(data []byte) error {
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

type GetResultOperatorBasicInfoList []GetResultOperatorBasicInfo

func (list *GetResultOperatorBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultOperatorBasicInfo)
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

type GetResultOperKeyWordList []goaliyun.String

func (list *GetResultOperKeyWordList) UnmarshalJSON(data []byte) error {
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

type GetResultReferenceList []goaliyun.String

func (list *GetResultReferenceList) UnmarshalJSON(data []byte) error {
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
