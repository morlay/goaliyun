package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SetVServerGroupAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
	VServerGroupName     string `position:"Query" name:"VServerGroupName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetVServerGroupAttributeRequest) Invoke(client goaliyun.Client) (*SetVServerGroupAttributeResponse, error) {
	resp := &SetVServerGroupAttributeResponse{}
	err := client.Request("slb", "SetVServerGroupAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetVServerGroupAttributeResponse struct {
	RequestId        goaliyun.String
	VServerGroupId   goaliyun.String
	VServerGroupName goaliyun.String
	BackendServers   SetVServerGroupAttributeBackendServerList
}

type SetVServerGroupAttributeBackendServer struct {
	ServerId goaliyun.String
	Port     goaliyun.Integer
	Weight   goaliyun.Integer
	Type     goaliyun.String
	ServerIp goaliyun.String
	VpcId    goaliyun.String
}

type SetVServerGroupAttributeBackendServerList []SetVServerGroupAttributeBackendServer

func (list *SetVServerGroupAttributeBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SetVServerGroupAttributeBackendServer)
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
