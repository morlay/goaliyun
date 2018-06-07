package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVServerGroupAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVServerGroupAttributeRequest) Invoke(client goaliyun.Client) (*DescribeVServerGroupAttributeResponse, error) {
	resp := &DescribeVServerGroupAttributeResponse{}
	err := client.Request("slb", "DescribeVServerGroupAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVServerGroupAttributeResponse struct {
	RequestId        goaliyun.String
	VServerGroupId   goaliyun.String
	VServerGroupName goaliyun.String
	BackendServers   DescribeVServerGroupAttributeBackendServerList
}

type DescribeVServerGroupAttributeBackendServer struct {
	ServerId goaliyun.String
	Port     goaliyun.Integer
	Weight   goaliyun.Integer
	Type     goaliyun.String
	ServerIp goaliyun.String
	VpcId    goaliyun.String
}

type DescribeVServerGroupAttributeBackendServerList []DescribeVServerGroupAttributeBackendServer

func (list *DescribeVServerGroupAttributeBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVServerGroupAttributeBackendServer)
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
