package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type PickLocalNumberRequest struct {
	InstanceId       string                              `position:"Query" name:"InstanceId"`
	CandidateNumbers *PickLocalNumberCandidateNumberList `position:"Query" type:"Repeated" name:"CandidateNumber"`
	CalleeNumber     string                              `position:"Query" name:"CalleeNumber"`
	RegionId         string                              `position:"Query" name:"RegionId"`
}

func (req *PickLocalNumberRequest) Invoke(client goaliyun.Client) (*PickLocalNumberResponse, error) {
	resp := &PickLocalNumberResponse{}
	err := client.Request("ccc", "PickLocalNumber", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PickLocalNumberResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      PickLocalNumberData
}

type PickLocalNumberData struct {
	Callee PickLocalNumberCallee
	Caller PickLocalNumberCaller
}

type PickLocalNumberCallee struct {
	Number   goaliyun.String
	Province goaliyun.String
	City     goaliyun.String
}

type PickLocalNumberCaller struct {
	Number   goaliyun.String
	Province goaliyun.String
	City     goaliyun.String
}

type PickLocalNumberCandidateNumberList []string

func (list *PickLocalNumberCandidateNumberList) UnmarshalJSON(data []byte) error {
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
