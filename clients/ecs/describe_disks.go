package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDisksRequest struct {
	Tag4Value                     string                                 `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId               int64                                  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId                    string                                 `position:"Query" name:"SnapshotId"`
	Tag2Key                       string                                 `position:"Query" name:"Tag.2.Key"`
	Filter2Value                  string                                 `position:"Query" name:"Filter.2.Value"`
	AutoSnapshotPolicyId          string                                 `position:"Query" name:"AutoSnapshotPolicyId"`
	Tag3Key                       string                                 `position:"Query" name:"Tag.3.Key"`
	PageNumber                    int64                                  `position:"Query" name:"PageNumber"`
	DiskName                      string                                 `position:"Query" name:"DiskName"`
	Tag1Value                     string                                 `position:"Query" name:"Tag.1.Value"`
	DeleteAutoSnapshot            string                                 `position:"Query" name:"DeleteAutoSnapshot"`
	ResourceGroupId               string                                 `position:"Query" name:"ResourceGroupId"`
	DiskChargeType                string                                 `position:"Query" name:"DiskChargeType"`
	LockReason                    string                                 `position:"Query" name:"LockReason"`
	Filter1Key                    string                                 `position:"Query" name:"Filter.1.Key"`
	PageSize                      int64                                  `position:"Query" name:"PageSize"`
	DiskIds                       string                                 `position:"Query" name:"DiskIds"`
	DeleteWithInstance            string                                 `position:"Query" name:"DeleteWithInstance"`
	Tag3Value                     string                                 `position:"Query" name:"Tag.3.Value"`
	EnableAutoSnapshot            string                                 `position:"Query" name:"EnableAutoSnapshot"`
	DryRun                        string                                 `position:"Query" name:"DryRun"`
	Tag5Key                       string                                 `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount          string                                 `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string                                 `position:"Query" name:"OwnerAccount"`
	Filter1Value                  string                                 `position:"Query" name:"Filter.1.Value"`
	Portable                      string                                 `position:"Query" name:"Portable"`
	EnableAutomatedSnapshotPolicy string                                 `position:"Query" name:"EnableAutomatedSnapshotPolicy"`
	Filter2Key                    string                                 `position:"Query" name:"Filter.2.Key"`
	OwnerId                       int64                                  `position:"Query" name:"OwnerId"`
	DiskType                      string                                 `position:"Query" name:"DiskType"`
	Tag5Value                     string                                 `position:"Query" name:"Tag.5.Value"`
	Tag1Key                       string                                 `position:"Query" name:"Tag.1.Key"`
	AdditionalAttributess         *DescribeDisksAdditionalAttributesList `position:"Query" type:"Repeated" name:"AdditionalAttributes"`
	EnableShared                  string                                 `position:"Query" name:"EnableShared"`
	InstanceId                    string                                 `position:"Query" name:"InstanceId"`
	Encrypted                     string                                 `position:"Query" name:"Encrypted"`
	Tag2Value                     string                                 `position:"Query" name:"Tag.2.Value"`
	ZoneId                        string                                 `position:"Query" name:"ZoneId"`
	Tag4Key                       string                                 `position:"Query" name:"Tag.4.Key"`
	Category                      string                                 `position:"Query" name:"Category"`
	Status                        string                                 `position:"Query" name:"Status"`
	RegionId                      string                                 `position:"Query" name:"RegionId"`
}

func (req *DescribeDisksRequest) Invoke(client goaliyun.Client) (*DescribeDisksResponse, error) {
	resp := &DescribeDisksResponse{}
	err := client.Request("ecs", "DescribeDisks", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDisksResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Disks      DescribeDisksDiskList
}

type DescribeDisksDisk struct {
	DiskId                        goaliyun.String
	RegionId                      goaliyun.String
	ZoneId                        goaliyun.String
	DiskName                      goaliyun.String
	Description                   goaliyun.String
	Type                          goaliyun.String
	Category                      goaliyun.String
	Size                          goaliyun.Integer
	ImageId                       goaliyun.String
	SourceSnapshotId              goaliyun.String
	AutoSnapshotPolicyId          goaliyun.String
	ProductCode                   goaliyun.String
	Portable                      bool
	Status                        goaliyun.String
	InstanceId                    goaliyun.String
	Device                        goaliyun.String
	DeleteWithInstance            bool
	DeleteAutoSnapshot            bool
	EnableAutoSnapshot            bool
	EnableAutomatedSnapshotPolicy bool
	CreationTime                  goaliyun.String
	AttachedTime                  goaliyun.String
	DetachedTime                  goaliyun.String
	DiskChargeType                goaliyun.String
	ExpiredTime                   goaliyun.String
	ResourceGroupId               goaliyun.String
	Encrypted                     bool
	MountInstanceNum              goaliyun.Integer
	IOPS                          goaliyun.Integer
	IOPSRead                      goaliyun.Integer
	IOPSWrite                     goaliyun.Integer
	OperationLocks                DescribeDisksOperationLockList
	MountInstances                DescribeDisksMountInstanceList
	Tags                          DescribeDisksTagList
}

type DescribeDisksOperationLock struct {
	LockReason goaliyun.String
}

type DescribeDisksMountInstance struct {
	InstanceId   goaliyun.String
	Device       goaliyun.String
	AttachedTime goaliyun.String
}

type DescribeDisksTag struct {
	TagKey   goaliyun.String
	TagValue goaliyun.String
}

type DescribeDisksAdditionalAttributesList []string

func (list *DescribeDisksAdditionalAttributesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeDisksDiskList []DescribeDisksDisk

func (list *DescribeDisksDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksDisk)
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

type DescribeDisksOperationLockList []DescribeDisksOperationLock

func (list *DescribeDisksOperationLockList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksOperationLock)
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

type DescribeDisksMountInstanceList []DescribeDisksMountInstance

func (list *DescribeDisksMountInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksMountInstance)
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

type DescribeDisksTagList []DescribeDisksTag

func (list *DescribeDisksTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksTag)
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
