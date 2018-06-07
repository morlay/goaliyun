package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFlowClusterHostRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListFlowClusterHostRequest) Invoke(client goaliyun.Client) (*ListFlowClusterHostResponse, error) {
	resp := &ListFlowClusterHostResponse{}
	err := client.Request("emr", "ListFlowClusterHost", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFlowClusterHostResponse struct {
	RequestId goaliyun.String
	HostList  ListFlowClusterHostHostList
}

type ListFlowClusterHostHost struct {
	HostId         goaliyun.String
	HostName       goaliyun.String
	PublicIp       goaliyun.String
	PrivateIp      goaliyun.String
	Role           goaliyun.String
	InstanceType   goaliyun.String
	Cpu            goaliyun.Integer
	Memory         goaliyun.Integer
	Status         goaliyun.String
	Type           goaliyun.String
	HostInstanceId goaliyun.String
	SerialNumber   goaliyun.String
}

type ListFlowClusterHostHostList []ListFlowClusterHostHost

func (list *ListFlowClusterHostHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowClusterHostHost)
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
