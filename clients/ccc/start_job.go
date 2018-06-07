package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type StartJobRequest struct {
	JobJson              string                     `position:"Query" name:"JobJson"`
	CallingNumbers       *StartJobCallingNumberList `position:"Query" type:"Repeated" name:"CallingNumber"`
	InstanceId           string                     `position:"Query" name:"InstanceId"`
	GroupId              string                     `position:"Query" name:"GroupId"`
	SelfHostedCallCenter string                     `position:"Query" name:"SelfHostedCallCenter"`
	ScenarioId           string                     `position:"Query" name:"ScenarioId"`
	RegionId             string                     `position:"Query" name:"RegionId"`
}

func (req *StartJobRequest) Invoke(client goaliyun.Client) (*StartJobResponse, error) {
	resp := &StartJobResponse{}
	err := client.Request("ccc", "StartJob", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartJobResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	TaskIds        StartJobKeyValuePairList
}

type StartJobKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type StartJobCallingNumberList []string

func (list *StartJobCallingNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type StartJobKeyValuePairList []StartJobKeyValuePair

func (list *StartJobKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartJobKeyValuePair)
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
