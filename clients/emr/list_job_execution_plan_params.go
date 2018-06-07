package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobExecutionPlanParamsRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RelateId        string `position:"Query" name:"RelateId"`
	ParamBizType    string `position:"Query" name:"ParamBizType"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListJobExecutionPlanParamsRequest) Invoke(client goaliyun.Client) (*ListJobExecutionPlanParamsResponse, error) {
	resp := &ListJobExecutionPlanParamsResponse{}
	err := client.Request("emr", "ListJobExecutionPlanParams", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobExecutionPlanParamsResponse struct {
	RequestId     goaliyun.String
	Success       goaliyun.String
	ErrCode       goaliyun.String
	ErrMsg        goaliyun.String
	ParamInfoList ListJobExecutionPlanParamsParamInfoList
}

type ListJobExecutionPlanParamsParamInfo struct {
	ParamBizType         goaliyun.String
	RelateId             goaliyun.String
	ParamName            goaliyun.String
	ParamValue           goaliyun.String
	UtcCreateTimestamp   goaliyun.Integer
	UtcModifiedTimestamp goaliyun.Integer
}

type ListJobExecutionPlanParamsParamInfoList []ListJobExecutionPlanParamsParamInfo

func (list *ListJobExecutionPlanParamsParamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExecutionPlanParamsParamInfo)
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
