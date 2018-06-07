package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateJobExecutionPlanParamRequest struct {
	ResourceOwnerId int64                                         `position:"Query" name:"ResourceOwnerId"`
	RelateId        string                                        `position:"Query" name:"RelateId"`
	WorkParamPairs  *CreateJobExecutionPlanParamWorkParamPairList `position:"Query" type:"Repeated" name:"WorkParamPair"`
	ParamBizType    string                                        `position:"Query" name:"ParamBizType"`
	RegionId        string                                        `position:"Query" name:"RegionId"`
}

func (req *CreateJobExecutionPlanParamRequest) Invoke(client goaliyun.Client) (*CreateJobExecutionPlanParamResponse, error) {
	resp := &CreateJobExecutionPlanParamResponse{}
	err := client.Request("emr", "CreateJobExecutionPlanParam", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateJobExecutionPlanParamWorkParamPair struct {
	Name  string `name:"Name"`
	Value string `name:"Value"`
}

type CreateJobExecutionPlanParamResponse struct {
	RequestId goaliyun.String
	Success   goaliyun.String
	ErrCode   goaliyun.String
	ErrMsg    goaliyun.String
	Ids       CreateJobExecutionPlanParamIdList
}

type CreateJobExecutionPlanParamWorkParamPairList []CreateJobExecutionPlanParamWorkParamPair

func (list *CreateJobExecutionPlanParamWorkParamPairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateJobExecutionPlanParamWorkParamPair)
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

type CreateJobExecutionPlanParamIdList []goaliyun.String

func (list *CreateJobExecutionPlanParamIdList) UnmarshalJSON(data []byte) error {
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
