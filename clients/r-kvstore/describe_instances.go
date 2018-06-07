package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceStatus       string `position:"Query" name:"InstanceStatus"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SearchKey            string `position:"Query" name:"SearchKey"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	ChargeType           string `position:"Query" name:"ChargeType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstancesRequest) Invoke(client goaliyun.Client) (*DescribeInstancesResponse, error) {
	resp := &DescribeInstancesResponse{}
	err := client.Request("r-kvstore", "DescribeInstances", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstancesResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	Instances  DescribeInstancesKVStoreInstanceList
}

type DescribeInstancesKVStoreInstance struct {
	ReplacateId         goaliyun.String
	InstanceId          goaliyun.String
	InstanceName        goaliyun.String
	ConnectionDomain    goaliyun.String
	Port                goaliyun.Integer
	UserName            goaliyun.String
	InstanceStatus      goaliyun.String
	RegionId            goaliyun.String
	Capacity            goaliyun.Integer
	InstanceClass       goaliyun.String
	QPS                 goaliyun.Integer
	Bandwidth           goaliyun.Integer
	Connections         goaliyun.Integer
	ZoneId              goaliyun.String
	Config              goaliyun.String
	ChargeType          goaliyun.String
	NetworkType         goaliyun.String
	VpcId               goaliyun.String
	VSwitchId           goaliyun.String
	PrivateIp           goaliyun.String
	CreateTime          goaliyun.String
	EndTime             goaliyun.String
	HasRenewChangeOrder goaliyun.String
	IsRds               bool
	InstanceType        goaliyun.String
	ArchitectureType    goaliyun.String
	NodeType            goaliyun.String
	PackageType         goaliyun.String
	EngineVersion       goaliyun.String
}

type DescribeInstancesKVStoreInstanceList []DescribeInstancesKVStoreInstance

func (list *DescribeInstancesKVStoreInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesKVStoreInstance)
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
