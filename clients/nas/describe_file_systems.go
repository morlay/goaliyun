package nas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeFileSystemsRequest struct {
	PageSize     int64  `position:"Query" name:"PageSize"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	FileSystemId string `position:"Query" name:"FileSystemId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeFileSystemsRequest) Invoke(client goaliyun.Client) (*DescribeFileSystemsResponse, error) {
	resp := &DescribeFileSystemsResponse{}
	err := client.Request("nas", "DescribeFileSystems", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFileSystemsResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageSize    goaliyun.Integer
	PageNumber  goaliyun.Integer
	FileSystems DescribeFileSystemsFileSystemList
}

type DescribeFileSystemsFileSystem struct {
	FileSystemId goaliyun.String
	Destription  goaliyun.String
	CreateTime   goaliyun.String
	RegionId     goaliyun.String
	ProtocolType goaliyun.String
	StorageType  goaliyun.String
	MeteredSize  goaliyun.Integer
	MountTargets DescribeFileSystemsMountTargetList
	Packages     DescribeFileSystems_PackageList
}

type DescribeFileSystemsMountTarget struct {
	MountTargetDomain goaliyun.String
}

type DescribeFileSystems_Package struct {
	PackageId goaliyun.String
}

type DescribeFileSystemsFileSystemList []DescribeFileSystemsFileSystem

func (list *DescribeFileSystemsFileSystemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFileSystemsFileSystem)
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

type DescribeFileSystemsMountTargetList []DescribeFileSystemsMountTarget

func (list *DescribeFileSystemsMountTargetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFileSystemsMountTarget)
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

type DescribeFileSystems_PackageList []DescribeFileSystems_Package

func (list *DescribeFileSystems_PackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFileSystems_Package)
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
