package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClustersRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListClustersRequest) Invoke(client goaliyun.Client) (*ListClustersResponse, error) {
	resp := &ListClustersResponse{}
	err := client.Request("ehpc", "ListClusters", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClustersResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Clusters   ListClustersClusterInfoSimpleList
}

type ListClustersClusterInfoSimple struct {
	Id              goaliyun.String
	RegionId        goaliyun.String
	ZoneId          goaliyun.String
	Name            goaliyun.String
	Description     goaliyun.String
	Status          goaliyun.String
	OsTag           goaliyun.String
	AccountType     goaliyun.String
	SchedulerType   goaliyun.String
	DeployMode      goaliyun.String
	Count           goaliyun.Integer
	InstanceType    goaliyun.String
	LoginNodes      goaliyun.String
	CreateTime      goaliyun.String
	ImageOwnerAlias goaliyun.String
	ImageId         goaliyun.String
	Managers        ListClustersManagers
	Computes        ListClustersComputes
	TotalResources  ListClustersTotalResources
	UsedResources   ListClustersUsedResources
}

type ListClustersManagers struct {
	Total          goaliyun.Integer
	NormalCount    goaliyun.Integer
	OperatingCount goaliyun.Integer
	StoppedCount   goaliyun.Integer
	ExceptionCount goaliyun.Integer
}

type ListClustersComputes struct {
	Total          goaliyun.Integer
	NormalCount    goaliyun.Integer
	OperatingCount goaliyun.Integer
	StoppedCount   goaliyun.Integer
	ExceptionCount goaliyun.Integer
}

type ListClustersTotalResources struct {
	Cpu    goaliyun.Integer
	Memory goaliyun.Integer
	Gpu    goaliyun.Integer
}

type ListClustersUsedResources struct {
	Cpu    goaliyun.Integer
	Memory goaliyun.Integer
	Gpu    goaliyun.Integer
}

type ListClustersClusterInfoSimpleList []ListClustersClusterInfoSimple

func (list *ListClustersClusterInfoSimpleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClustersClusterInfoSimple)
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
