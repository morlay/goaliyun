package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateMasterSlaveServerGroupRequest struct {
	Access_key_id              string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveBackendServers  string `position:"Query" name:"MasterSlaveBackendServers"`
	LoadBalancerId             string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	MasterSlaveServerGroupName string `position:"Query" name:"MasterSlaveServerGroupName"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	Tags                       string `position:"Query" name:"Tags"`
	RegionId                   string `position:"Query" name:"RegionId"`
}

func (req *CreateMasterSlaveServerGroupRequest) Invoke(client goaliyun.Client) (*CreateMasterSlaveServerGroupResponse, error) {
	resp := &CreateMasterSlaveServerGroupResponse{}
	err := client.Request("slb", "CreateMasterSlaveServerGroup", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateMasterSlaveServerGroupResponse struct {
	RequestId                 goaliyun.String
	MasterSlaveServerGroupId  goaliyun.String
	MasterSlaveBackendServers CreateMasterSlaveServerGroupMasterSlaveBackendServerList
}

type CreateMasterSlaveServerGroupMasterSlaveBackendServer struct {
	ServerId   goaliyun.String
	Port       goaliyun.Integer
	Weight     goaliyun.Integer
	ServerType goaliyun.String
	Type       goaliyun.String
	ServerIp   goaliyun.String
	VpcId      goaliyun.String
}

type CreateMasterSlaveServerGroupMasterSlaveBackendServerList []CreateMasterSlaveServerGroupMasterSlaveBackendServer

func (list *CreateMasterSlaveServerGroupMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateMasterSlaveServerGroupMasterSlaveBackendServer)
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
