package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddBackendServersRequest struct {
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

func (req *AddBackendServersRequest) Invoke(client goaliyun.Client) (*AddBackendServersResponse, error) {
	resp := &AddBackendServersResponse{}
	err := client.Request("slb", "AddBackendServers", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddBackendServersResponse struct {
	RequestId      goaliyun.String
	LoadBalancerId goaliyun.String
	BackendServers AddBackendServersBackendServerList
}

type AddBackendServersBackendServer struct {
	ServerId goaliyun.String
	Weight   goaliyun.String
	ServerIp goaliyun.String
	VpcId    goaliyun.String
	Type     goaliyun.String
}

type AddBackendServersBackendServerList []AddBackendServersBackendServer

func (list *AddBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddBackendServersBackendServer)
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
