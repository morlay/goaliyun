package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListMyGroupInstancesRequest struct {
	Total       string `position:"Query" name:"Total"`
	InstanceIds string `position:"Query" name:"InstanceIds"`
	GroupId     int64  `position:"Query" name:"GroupId"`
	PageSize    int64  `position:"Query" name:"PageSize"`
	Category    string `position:"Query" name:"Category"`
	Keyword     string `position:"Query" name:"Keyword"`
	PageNumber  int64  `position:"Query" name:"PageNumber"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ListMyGroupInstancesRequest) Invoke(client goaliyun.Client) (*ListMyGroupInstancesResponse, error) {
	resp := &ListMyGroupInstancesResponse{}
	err := client.Request("cms", "ListMyGroupInstances", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListMyGroupInstancesResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	Total        goaliyun.Integer
	Resources    ListMyGroupInstancesResourceList
}

type ListMyGroupInstancesResource struct {
	Id           goaliyun.Integer
	AliUid       goaliyun.Integer
	RegionId     goaliyun.String
	InstanceId   goaliyun.String
	Category     goaliyun.String
	InstanceName goaliyun.String
}

type ListMyGroupInstancesResourceList []ListMyGroupInstancesResource

func (list *ListMyGroupInstancesResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupInstancesResource)
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
