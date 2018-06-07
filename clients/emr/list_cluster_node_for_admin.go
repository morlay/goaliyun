package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterNodeForAdminRequest struct {
	ResourceOwnerId int64                                  `position:"Query" name:"ResourceOwnerId"`
	StatusLists     *ListClusterNodeForAdminStatusListList `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize        int64                                  `position:"Query" name:"PageSize"`
	ClusterId       string                                 `position:"Query" name:"ClusterId"`
	UserId          string                                 `position:"Query" name:"UserId"`
	PageNumber      int64                                  `position:"Query" name:"PageNumber"`
	RegionId        string                                 `position:"Query" name:"RegionId"`
}

func (req *ListClusterNodeForAdminRequest) Invoke(client goaliyun.Client) (*ListClusterNodeForAdminResponse, error) {
	resp := &ListClusterNodeForAdminResponse{}
	err := client.Request("emr", "ListClusterNodeForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterNodeForAdminResponse struct {
	RequestId       goaliyun.String
	ClusterNodeList ListClusterNodeForAdminClusterNodeList
}

type ListClusterNodeForAdminClusterNode struct {
	ClusterId       goaliyun.String
	CpuCore         goaliyun.String
	Daemons         goaliyun.String
	DiskDevices     goaliyun.String
	DiskInfo        goaliyun.String
	DiskType        goaliyun.String
	GmtCreate       goaliyun.String
	GmtModified     goaliyun.String
	HostName        goaliyun.String
	Id              goaliyun.String
	ImageId         goaliyun.String
	InnerIpAddress  goaliyun.String
	InstanceId      goaliyun.String
	InstanceType    goaliyun.String
	IsMaster        goaliyun.String
	MemCapacity     goaliyun.String
	Payment         goaliyun.String
	PublicIpAddress goaliyun.String
	RegionId        goaliyun.String
	SecurityGroupId goaliyun.String
	SerialNumber    goaliyun.String
	Status          goaliyun.String
	UserId          goaliyun.String
	ZoneId          goaliyun.String
}

type ListClusterNodeForAdminStatusListList []string

func (list *ListClusterNodeForAdminStatusListList) UnmarshalJSON(data []byte) error {
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

type ListClusterNodeForAdminClusterNodeList []ListClusterNodeForAdminClusterNode

func (list *ListClusterNodeForAdminClusterNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterNodeForAdminClusterNode)
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
