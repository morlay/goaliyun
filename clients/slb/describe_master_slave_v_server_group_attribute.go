package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMasterSlaveVServerGroupAttributeRequest struct {
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveVServerGroupId string `position:"Query" name:"MasterSlaveVServerGroupId"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
	RegionId                  string `position:"Query" name:"RegionId"`
}

func (req *DescribeMasterSlaveVServerGroupAttributeRequest) Invoke(client goaliyun.Client) (*DescribeMasterSlaveVServerGroupAttributeResponse, error) {
	resp := &DescribeMasterSlaveVServerGroupAttributeResponse{}
	err := client.Request("slb", "DescribeMasterSlaveVServerGroupAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMasterSlaveVServerGroupAttributeResponse struct {
	RequestId                   goaliyun.String
	MasterSlaveVServerGroupId   goaliyun.String
	MasterSlaveVServerGroupName goaliyun.String
	MasterSlaveBackendServers   DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServerList
}

type DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServer struct {
	ServerId goaliyun.String
	Port     goaliyun.Integer
	Weight   goaliyun.Integer
	IsBackup goaliyun.Integer
	Type     goaliyun.String
	ServerIp goaliyun.String
	VpcId    goaliyun.String
}

type DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServerList []DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServer

func (list *DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServer)
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
