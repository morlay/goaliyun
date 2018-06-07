package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterHostRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostName        string `position:"Query" name:"HostName"`
	HostInstanceId  string `position:"Query" name:"HostInstanceId"`
	PrivateIp       string `position:"Query" name:"PrivateIp"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ComponentName   string `position:"Query" name:"ComponentName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterHostRequest) Invoke(client goaliyun.Client) (*ListClusterHostResponse, error) {
	resp := &ListClusterHostResponse{}
	err := client.Request("emr", "ListClusterHost", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterHostResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Total      goaliyun.Integer
	HostList   ListClusterHostHostList
}

type ListClusterHostHost struct {
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

type ListClusterHostHostList []ListClusterHostHost

func (list *ListClusterHostHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHostHost)
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
