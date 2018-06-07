package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListContactFlowsRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListContactFlowsRequest) Invoke(client goaliyun.Client) (*ListContactFlowsResponse, error) {
	resp := &ListContactFlowsResponse{}
	err := client.Request("ccc", "ListContactFlows", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListContactFlowsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	ContactFlows   ListContactFlowsContactFlowList
}

type ListContactFlowsContactFlow struct {
	ContactFlowId          goaliyun.String
	InstanceId             goaliyun.String
	ContactFlowName        goaliyun.String
	ContactFlowDescription goaliyun.String
	Type                   goaliyun.String
	AppliedVersion         goaliyun.String
	Versions               ListContactFlowsContactFlowVersionList
	PhoneNumbers           ListContactFlowsPhoneNumberList
}

type ListContactFlowsContactFlowVersion struct {
	ContactFlowVersionId          goaliyun.String
	Version                       goaliyun.String
	ContactFlowVersionDescription goaliyun.String
	LastModified                  goaliyun.String
	LastModifiedBy                goaliyun.String
	LockedBy                      goaliyun.String
	Status                        goaliyun.String
}

type ListContactFlowsPhoneNumber struct {
	PhoneNumberId          goaliyun.String
	InstanceId             goaliyun.String
	Number                 goaliyun.String
	PhoneNumberDescription goaliyun.String
	TestOnly               bool
	RemainingTime          goaliyun.Integer
	AllowOutbound          bool
	Usage                  goaliyun.String
	Trunks                 goaliyun.Integer
}

type ListContactFlowsContactFlowList []ListContactFlowsContactFlow

func (list *ListContactFlowsContactFlowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListContactFlowsContactFlow)
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

type ListContactFlowsContactFlowVersionList []ListContactFlowsContactFlowVersion

func (list *ListContactFlowsContactFlowVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListContactFlowsContactFlowVersion)
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

type ListContactFlowsPhoneNumberList []ListContactFlowsPhoneNumber

func (list *ListContactFlowsPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListContactFlowsPhoneNumber)
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
