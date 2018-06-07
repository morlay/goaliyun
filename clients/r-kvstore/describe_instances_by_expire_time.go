package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstancesByExpireTimeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	HasExpiredRes        string `position:"Query" name:"HasExpiredRes"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	ExpirePeriod         int64  `position:"Query" name:"ExpirePeriod"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstancesByExpireTimeRequest) Invoke(client goaliyun.Client) (*DescribeInstancesByExpireTimeResponse, error) {
	resp := &DescribeInstancesByExpireTimeResponse{}
	err := client.Request("r-kvstore", "DescribeInstancesByExpireTime", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstancesByExpireTimeResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	Instances  DescribeInstancesByExpireTimeKVStoreInstanceList
}

type DescribeInstancesByExpireTimeKVStoreInstance struct {
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

type DescribeInstancesByExpireTimeKVStoreInstanceList []DescribeInstancesByExpireTimeKVStoreInstance

func (list *DescribeInstancesByExpireTimeKVStoreInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesByExpireTimeKVStoreInstance)
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
