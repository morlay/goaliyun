package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeGlobalAccelerationInstancesRequest struct {
	IpAddress                    string `position:"Query" name:"IpAddress"`
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthType                string `position:"Query" name:"BandwidthType"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	ServiceLocation              string `position:"Query" name:"ServiceLocation"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	ServerId                     string `position:"Query" name:"ServerId"`
	PageNumber                   int64  `position:"Query" name:"PageNumber"`
	Name                         string `position:"Query" name:"Name"`
	PageSize                     int64  `position:"Query" name:"PageSize"`
	Status                       string `position:"Query" name:"Status"`
	RegionId                     string `position:"Query" name:"RegionId"`
}

func (req *DescribeGlobalAccelerationInstancesRequest) Invoke(client goaliyun.Client) (*DescribeGlobalAccelerationInstancesResponse, error) {
	resp := &DescribeGlobalAccelerationInstancesResponse{}
	err := client.Request("vpc", "DescribeGlobalAccelerationInstances", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeGlobalAccelerationInstancesResponse struct {
	RequestId                   goaliyun.String
	TotalCount                  goaliyun.Integer
	PageNumber                  goaliyun.Integer
	PageSize                    goaliyun.Integer
	GlobalAccelerationInstances DescribeGlobalAccelerationInstancesGlobalAccelerationInstanceList
}

type DescribeGlobalAccelerationInstancesGlobalAccelerationInstance struct {
	RegionId                     goaliyun.String
	GlobalAccelerationInstanceId goaliyun.String
	IpAddress                    goaliyun.String
	Status                       goaliyun.String
	Bandwidth                    goaliyun.String
	InternetChargeType           goaliyun.String
	ChargeType                   goaliyun.String
	BandwidthType                goaliyun.String
	AccelerationLocation         goaliyun.String
	ServiceLocation              goaliyun.String
	Name                         goaliyun.String
	Description                  goaliyun.String
	ExpiredTime                  goaliyun.String
	CreationTime                 goaliyun.String
	OperationLocks               DescribeGlobalAccelerationInstancesLockReasonList
	BackendServers               DescribeGlobalAccelerationInstancesBackendServerList
	PublicIpAddresses            DescribeGlobalAccelerationInstancesPublicIpAddressList
}

type DescribeGlobalAccelerationInstancesLockReason struct {
	LockReason goaliyun.String
}

type DescribeGlobalAccelerationInstancesBackendServer struct {
	RegionId        goaliyun.String
	ServerId        goaliyun.String
	ServerIpAddress goaliyun.String
	ServerType      goaliyun.String
}

type DescribeGlobalAccelerationInstancesPublicIpAddress struct {
	AllocationId goaliyun.String
	IpAddress    goaliyun.String
}

type DescribeGlobalAccelerationInstancesGlobalAccelerationInstanceList []DescribeGlobalAccelerationInstancesGlobalAccelerationInstance

func (list *DescribeGlobalAccelerationInstancesGlobalAccelerationInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesGlobalAccelerationInstance)
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

type DescribeGlobalAccelerationInstancesLockReasonList []DescribeGlobalAccelerationInstancesLockReason

func (list *DescribeGlobalAccelerationInstancesLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesLockReason)
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

type DescribeGlobalAccelerationInstancesBackendServerList []DescribeGlobalAccelerationInstancesBackendServer

func (list *DescribeGlobalAccelerationInstancesBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesBackendServer)
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

type DescribeGlobalAccelerationInstancesPublicIpAddressList []DescribeGlobalAccelerationInstancesPublicIpAddress

func (list *DescribeGlobalAccelerationInstancesPublicIpAddressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesPublicIpAddress)
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
