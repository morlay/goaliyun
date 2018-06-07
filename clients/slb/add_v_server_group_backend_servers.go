package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddVServerGroupBackendServersRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddVServerGroupBackendServersRequest) Invoke(client goaliyun.Client) (*AddVServerGroupBackendServersResponse, error) {
	resp := &AddVServerGroupBackendServersResponse{}
	err := client.Request("slb", "AddVServerGroupBackendServers", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddVServerGroupBackendServersResponse struct {
	RequestId      goaliyun.String
	VServerGroupId goaliyun.String
	BackendServers AddVServerGroupBackendServersBackendServerList
}

type AddVServerGroupBackendServersBackendServer struct {
	ServerId goaliyun.String
	Port     goaliyun.Integer
	Weight   goaliyun.Integer
	Type     goaliyun.String
	ServerIp goaliyun.String
	VpcId    goaliyun.String
}

type AddVServerGroupBackendServersBackendServerList []AddVServerGroupBackendServersBackendServer

func (list *AddVServerGroupBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddVServerGroupBackendServersBackendServer)
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
