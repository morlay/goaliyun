package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SetBackendServersRequest struct {
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

func (req *SetBackendServersRequest) Invoke(client goaliyun.Client) (*SetBackendServersResponse, error) {
	resp := &SetBackendServersResponse{}
	err := client.Request("slb", "SetBackendServers", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetBackendServersResponse struct {
	RequestId      goaliyun.String
	LoadBalancerId goaliyun.String
	BackendServers SetBackendServersBackendServerList
}

type SetBackendServersBackendServer struct {
	ServerId goaliyun.String
	Weight   goaliyun.String
	ServerIp goaliyun.String
	VpcId    goaliyun.String
	Type     goaliyun.String
}

type SetBackendServersBackendServerList []SetBackendServersBackendServer

func (list *SetBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SetBackendServersBackendServer)
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
