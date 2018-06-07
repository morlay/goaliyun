package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterHealthRequest struct {
	ResourceOwnerId int64                               `position:"Query" name:"ResourceOwnerId"`
	ClusterIdLists  *ListClusterHealthClusterIdListList `position:"Query" type:"Repeated" name:"ClusterIdList"`
	RegionId        string                              `position:"Query" name:"RegionId"`
}

func (req *ListClusterHealthRequest) Invoke(client goaliyun.Client) (*ListClusterHealthResponse, error) {
	resp := &ListClusterHealthResponse{}
	err := client.Request("emr", "ListClusterHealth", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterHealthResponse struct {
	RequestId                 goaliyun.String
	ClusterHealthResponseList ListClusterHealthClusterHealthResponseList
}

type ListClusterHealthClusterHealthResponse struct {
	ClusterId             goaliyun.Integer
	ServiceHealthInfoList ListClusterHealthServiceHealthInfoList
	HealthResult          ListClusterHealthHealthResult
}

type ListClusterHealthServiceHealthInfo struct {
	Key           goaliyun.String
	PassNumber    goaliyun.Integer
	ErrorNumber   goaliyun.Integer
	WarningNumber goaliyun.Integer
	UnKnownNumber goaliyun.Integer
}

type ListClusterHealthHealthResult struct {
	Key           goaliyun.String
	PassNumber    goaliyun.Integer
	ErrorNumber   goaliyun.Integer
	WarningNumber goaliyun.Integer
	UnKnownNumber goaliyun.Integer
}

type ListClusterHealthClusterIdListList []string

func (list *ListClusterHealthClusterIdListList) UnmarshalJSON(data []byte) error {
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

type ListClusterHealthClusterHealthResponseList []ListClusterHealthClusterHealthResponse

func (list *ListClusterHealthClusterHealthResponseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHealthClusterHealthResponse)
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

type ListClusterHealthServiceHealthInfoList []ListClusterHealthServiceHealthInfo

func (list *ListClusterHealthServiceHealthInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHealthServiceHealthInfo)
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
