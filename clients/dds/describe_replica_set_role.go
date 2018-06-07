package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeReplicaSetRoleRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeReplicaSetRoleRequest) Invoke(client goaliyun.Client) (*DescribeReplicaSetRoleResponse, error) {
	resp := &DescribeReplicaSetRoleResponse{}
	err := client.Request("dds", "DescribeReplicaSetRole", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeReplicaSetRoleResponse struct {
	RequestId    goaliyun.String
	DBInstanceId goaliyun.String
	ReplicaSets  DescribeReplicaSetRoleReplicaSetList
}

type DescribeReplicaSetRoleReplicaSet struct {
	ReplicaSetRole   goaliyun.String
	ConnectionDomain goaliyun.String
	ConnectionPort   goaliyun.String
	ExpiredTime      goaliyun.String
	NetworkType      goaliyun.String
}

type DescribeReplicaSetRoleReplicaSetList []DescribeReplicaSetRoleReplicaSet

func (list *DescribeReplicaSetRoleReplicaSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaSetRoleReplicaSet)
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
