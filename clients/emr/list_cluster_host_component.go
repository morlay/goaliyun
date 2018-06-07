package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterHostComponentRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostName        string `position:"Query" name:"HostName"`
	HostInstanceId  string `position:"Query" name:"HostInstanceId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ComponentName   string `position:"Query" name:"ComponentName"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterHostComponentRequest) Invoke(client goaliyun.Client) (*ListClusterHostComponentResponse, error) {
	resp := &ListClusterHostComponentResponse{}
	err := client.Request("emr", "ListClusterHostComponent", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterHostComponentResponse struct {
	RequestId     goaliyun.String
	PageNumber    goaliyun.Integer
	PageSize      goaliyun.Integer
	Total         goaliyun.Integer
	ComponentList ListClusterHostComponentComponentList
}

type ListClusterHostComponentComponent struct {
	ServiceName          goaliyun.String
	ServiceDisplayName   goaliyun.String
	ComponentName        goaliyun.String
	ComponentDisplayName goaliyun.String
	Status               goaliyun.String
	NeedRestart          bool
	HostId               goaliyun.String
	ServerStatus         goaliyun.String
	HostName             goaliyun.String
	PublicIp             goaliyun.String
	PrivateIp            goaliyun.String
	Role                 goaliyun.String
	InstanceType         goaliyun.String
	Cpu                  goaliyun.Integer
	Memory               goaliyun.Integer
	HostInstanceId       goaliyun.String
	SerialNumber         goaliyun.String
}

type ListClusterHostComponentComponentList []ListClusterHostComponentComponent

func (list *ListClusterHostComponentComponentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHostComponentComponent)
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
