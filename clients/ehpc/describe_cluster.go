package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterRequest struct {
	ClusterId string `position:"Query" name:"ClusterId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterRequest) Invoke(client goaliyun.Client) (*DescribeClusterResponse, error) {
	resp := &DescribeClusterResponse{}
	err := client.Request("ehpc", "DescribeCluster", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterResponse struct {
	RequestId   goaliyun.String
	ClusterInfo DescribeClusterClusterInfo
}

type DescribeClusterClusterInfo struct {
	Id               goaliyun.String
	RegionId         goaliyun.String
	Name             goaliyun.String
	Description      goaliyun.String
	Status           goaliyun.String
	OsTag            goaliyun.String
	AccountType      goaliyun.String
	SchedulerType    goaliyun.String
	CreateTime       goaliyun.String
	SecurityGroupId  goaliyun.String
	VSwitchId        goaliyun.String
	VolumeType       goaliyun.String
	VolumeId         goaliyun.String
	VolumeProtocol   goaliyun.String
	VolumeMountpoint goaliyun.String
	RemoteDirectory  goaliyun.String
	DeployMode       goaliyun.String
	HaEnable         bool
	EcsChargeType    goaliyun.String
	KeyPairName      goaliyun.String
	SccClusterId     goaliyun.String
	ClientVersion    goaliyun.String
	ImageOwnerAlias  goaliyun.String
	ImageId          goaliyun.String
	Applications     DescribeClusterApplicationInfoList
	EcsInfo          DescribeClusterEcsInfo
}

type DescribeClusterApplicationInfo struct {
	Tag     goaliyun.String
	Name    goaliyun.String
	Version goaliyun.String
}

type DescribeClusterEcsInfo struct {
	Manager DescribeClusterManager
	Compute DescribeClusterCompute
	Login   DescribeClusterLogin
}

type DescribeClusterManager struct {
	Count        goaliyun.Integer
	InstanceType goaliyun.String
}

type DescribeClusterCompute struct {
	Count        goaliyun.Integer
	InstanceType goaliyun.String
}

type DescribeClusterLogin struct {
	Count        goaliyun.Integer
	InstanceType goaliyun.String
}

type DescribeClusterApplicationInfoList []DescribeClusterApplicationInfo

func (list *DescribeClusterApplicationInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterApplicationInfo)
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
