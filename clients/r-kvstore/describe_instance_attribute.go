package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstanceAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceAttributeRequest) Invoke(client goaliyun.Client) (*DescribeInstanceAttributeResponse, error) {
	resp := &DescribeInstanceAttributeResponse{}
	err := client.Request("r-kvstore", "DescribeInstanceAttribute", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceAttributeResponse struct {
	RequestId goaliyun.String
	Instances DescribeInstanceAttributeDBInstanceAttributeList
}

type DescribeInstanceAttributeDBInstanceAttribute struct {
	InstanceId          goaliyun.String
	InstanceName        goaliyun.String
	ConnectionDomain    goaliyun.String
	Port                goaliyun.Integer
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
	NodeType            goaliyun.String
	NetworkType         goaliyun.String
	VpcId               goaliyun.String
	VSwitchId           goaliyun.String
	PrivateIp           goaliyun.String
	CreateTime          goaliyun.String
	EndTime             goaliyun.String
	HasRenewChangeOrder goaliyun.String
	IsRds               bool
	Engine              goaliyun.String
	EngineVersion       goaliyun.String
	MaintainStartTime   goaliyun.String
	MaintainEndTime     goaliyun.String
	AvailabilityValue   goaliyun.String
	SecurityIPList      goaliyun.String
	InstanceType        goaliyun.String
	ArchitectureType    goaliyun.String
	NodeType1           goaliyun.String
	PackageType         goaliyun.String
	ReplacateId         goaliyun.String
	EngineVersion2      goaliyun.String
}

type DescribeInstanceAttributeDBInstanceAttributeList []DescribeInstanceAttributeDBInstanceAttribute

func (list *DescribeInstanceAttributeDBInstanceAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAttributeDBInstanceAttribute)
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
