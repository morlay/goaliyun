package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMasterSlaveServerGroupAttributeRequest struct {
	Access_key_id            string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveServerGroupId string `position:"Query" name:"MasterSlaveServerGroupId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	Tags                     string `position:"Query" name:"Tags"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *DescribeMasterSlaveServerGroupAttributeRequest) Invoke(client goaliyun.Client) (*DescribeMasterSlaveServerGroupAttributeResponse, error) {
	resp := &DescribeMasterSlaveServerGroupAttributeResponse{}
	err := client.Request("slb", "DescribeMasterSlaveServerGroupAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMasterSlaveServerGroupAttributeResponse struct {
	RequestId                  goaliyun.String
	MasterSlaveServerGroupId   goaliyun.String
	MasterSlaveServerGroupName goaliyun.String
	MasterSlaveBackendServers  DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList
}

type DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer struct {
	ServerId   goaliyun.String
	Port       goaliyun.Integer
	Weight     goaliyun.Integer
	ServerType goaliyun.String
	Type       goaliyun.String
	ServerIp   goaliyun.String
	VpcId      goaliyun.String
}

type DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList []DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer

func (list *DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer)
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
