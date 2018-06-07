package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListNodesRequest struct {
	HostName   string `position:"Query" name:"HostName"`
	Role       string `position:"Query" name:"Role"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	ClusterId  string `position:"Query" name:"ClusterId"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListNodesRequest) Invoke(client goaliyun.Client) (*ListNodesResponse, error) {
	resp := &ListNodesResponse{}
	err := client.Request("ehpc", "ListNodes", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListNodesResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Nodes      ListNodesNodeInfoList
}

type ListNodesNodeInfo struct {
	Id              goaliyun.String
	RegionId        goaliyun.String
	Status          goaliyun.String
	CreatedByEhpc   bool
	AddTime         goaliyun.String
	Expired         bool
	ExpiredTime     goaliyun.String
	SpotStrategy    goaliyun.String
	LockReason      goaliyun.String
	ImageOwnerAlias goaliyun.String
	ImageId         goaliyun.String
	Roles           ListNodesRoleList
	TotalResources  ListNodesTotalResources
	UsedResources   ListNodesUsedResources
}

type ListNodesTotalResources struct {
	Cpu    goaliyun.Integer
	Memory goaliyun.Integer
	Gpu    goaliyun.Integer
}

type ListNodesUsedResources struct {
	Cpu    goaliyun.Integer
	Memory goaliyun.Integer
	Gpu    goaliyun.Integer
}

type ListNodesNodeInfoList []ListNodesNodeInfo

func (list *ListNodesNodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListNodesNodeInfo)
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

type ListNodesRoleList []goaliyun.String

func (list *ListNodesRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
