package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RemoveBackendServersRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RemoveBackendServersRequest) Invoke(client goaliyun.Client) (*RemoveBackendServersResponse, error) {
	resp := &RemoveBackendServersResponse{}
	err := client.Request("slb", "RemoveBackendServers", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveBackendServersResponse struct {
	RequestId      goaliyun.String
	LoadBalancerId goaliyun.String
	BackendServers RemoveBackendServersBackendServerList
}

type RemoveBackendServersBackendServer struct {
	ServerId goaliyun.String
	Weight   goaliyun.Integer
	ServerIp goaliyun.String
	VpcId    goaliyun.String
	Type     goaliyun.String
}

type RemoveBackendServersBackendServerList []RemoveBackendServersBackendServer

func (list *RemoveBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveBackendServersBackendServer)
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
