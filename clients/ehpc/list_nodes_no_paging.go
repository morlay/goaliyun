package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListNodesNoPagingRequest struct {
	HostName     string `position:"Query" name:"HostName"`
	Role         string `position:"Query" name:"Role"`
	ClusterId    string `position:"Query" name:"ClusterId"`
	OnlyDetached string `position:"Query" name:"OnlyDetached"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ListNodesNoPagingRequest) Invoke(client goaliyun.Client) (*ListNodesNoPagingResponse, error) {
	resp := &ListNodesNoPagingResponse{}
	err := client.Request("ehpc", "ListNodesNoPaging", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListNodesNoPagingResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Nodes      ListNodesNoPagingNodeInfoList
}

type ListNodesNoPagingNodeInfo struct {
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
	Roles           ListNodesNoPagingRoleList
	TotalResources  ListNodesNoPagingTotalResources
	UsedResources   ListNodesNoPagingUsedResources
}

type ListNodesNoPagingTotalResources struct {
	Cpu    goaliyun.Integer
	Memory goaliyun.Integer
	Gpu    goaliyun.Integer
}

type ListNodesNoPagingUsedResources struct {
	Cpu    goaliyun.Integer
	Memory goaliyun.Integer
	Gpu    goaliyun.Integer
}

type ListNodesNoPagingNodeInfoList []ListNodesNoPagingNodeInfo

func (list *ListNodesNoPagingNodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListNodesNoPagingNodeInfo)
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

type ListNodesNoPagingRoleList []goaliyun.String

func (list *ListNodesNoPagingRoleList) UnmarshalJSON(data []byte) error {
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
