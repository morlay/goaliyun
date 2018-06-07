package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyVServerGroupBackendServersRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OldBackendServers    string `position:"Query" name:"OldBackendServers"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	NewBackendServers    string `position:"Query" name:"NewBackendServers"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyVServerGroupBackendServersRequest) Invoke(client goaliyun.Client) (*ModifyVServerGroupBackendServersResponse, error) {
	resp := &ModifyVServerGroupBackendServersResponse{}
	err := client.Request("slb", "ModifyVServerGroupBackendServers", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyVServerGroupBackendServersResponse struct {
	RequestId      goaliyun.String
	VServerGroupId goaliyun.String
	BackendServers ModifyVServerGroupBackendServersBackendServerList
}

type ModifyVServerGroupBackendServersBackendServer struct {
	ServerId goaliyun.String
	Port     goaliyun.Integer
	Weight   goaliyun.Integer
	Type     goaliyun.String
	ServerIp goaliyun.String
	VpcId    goaliyun.String
}

type ModifyVServerGroupBackendServersBackendServerList []ModifyVServerGroupBackendServersBackendServer

func (list *ModifyVServerGroupBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyVServerGroupBackendServersBackendServer)
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
