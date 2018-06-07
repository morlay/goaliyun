package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateMasterSlaveVServerGroupRequest struct {
	Access_key_id               string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId             int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveBackendServers   string `position:"Query" name:"MasterSlaveBackendServers"`
	LoadBalancerId              string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	MasterSlaveVServerGroupName string `position:"Query" name:"MasterSlaveVServerGroupName"`
	OwnerId                     int64  `position:"Query" name:"OwnerId"`
	Tags                        string `position:"Query" name:"Tags"`
	RegionId                    string `position:"Query" name:"RegionId"`
}

func (req *CreateMasterSlaveVServerGroupRequest) Invoke(client goaliyun.Client) (*CreateMasterSlaveVServerGroupResponse, error) {
	resp := &CreateMasterSlaveVServerGroupResponse{}
	err := client.Request("slb", "CreateMasterSlaveVServerGroup", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateMasterSlaveVServerGroupResponse struct {
	RequestId                 goaliyun.String
	MasterSlaveVServerGroupId goaliyun.String
	MasterSlaveBackendServers CreateMasterSlaveVServerGroupMasterSlaveBackendServerList
}

type CreateMasterSlaveVServerGroupMasterSlaveBackendServer struct {
	ServerId goaliyun.String
	Port     goaliyun.Integer
	Weight   goaliyun.Integer
	IsBackup goaliyun.Integer
	Type     goaliyun.String
	ServerIp goaliyun.String
	VpcId    goaliyun.String
}

type CreateMasterSlaveVServerGroupMasterSlaveBackendServerList []CreateMasterSlaveVServerGroupMasterSlaveBackendServer

func (list *CreateMasterSlaveVServerGroupMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateMasterSlaveVServerGroupMasterSlaveBackendServer)
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
