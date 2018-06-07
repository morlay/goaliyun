package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateVServerGroupRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
	VServerGroupName     string `position:"Query" name:"VServerGroupName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateVServerGroupRequest) Invoke(client goaliyun.Client) (*CreateVServerGroupResponse, error) {
	resp := &CreateVServerGroupResponse{}
	err := client.Request("slb", "CreateVServerGroup", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateVServerGroupResponse struct {
	RequestId      goaliyun.String
	VServerGroupId goaliyun.String
	BackendServers CreateVServerGroupBackendServerList
}

type CreateVServerGroupBackendServer struct {
	ServerId goaliyun.String
	Port     goaliyun.Integer
	Weight   goaliyun.Integer
	Type     goaliyun.String
	ServerIp goaliyun.String
	VpcId    goaliyun.String
}

type CreateVServerGroupBackendServerList []CreateVServerGroupBackendServer

func (list *CreateVServerGroupBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateVServerGroupBackendServer)
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
