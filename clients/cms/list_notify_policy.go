package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListNotifyPolicyRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	AlertName  string `position:"Query" name:"AlertName"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	Id         string `position:"Query" name:"Id"`
	Dimensions string `position:"Query" name:"Dimensions"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListNotifyPolicyRequest) Invoke(client goaliyun.Client) (*ListNotifyPolicyResponse, error) {
	resp := &ListNotifyPolicyResponse{}
	err := client.Request("cms", "ListNotifyPolicy", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListNotifyPolicyResponse struct {
	Code             goaliyun.String
	Message          goaliyun.String
	Success          goaliyun.String
	TraceId          goaliyun.String
	Total            goaliyun.Integer
	NotifyPolicyList ListNotifyPolicyNotifyPolicyList
}

type ListNotifyPolicyNotifyPolicy struct {
	AlertName  goaliyun.String
	Dimensions goaliyun.String
	Type       goaliyun.String
	Id         goaliyun.String
	StartTime  goaliyun.Integer
	EndTime    goaliyun.Integer
}

type ListNotifyPolicyNotifyPolicyList []ListNotifyPolicyNotifyPolicy

func (list *ListNotifyPolicyNotifyPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListNotifyPolicyNotifyPolicy)
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
