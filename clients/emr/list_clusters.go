package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClustersRequest struct {
	ResourceOwnerId  int64                            `position:"Query" name:"ResourceOwnerId"`
	StatusLists      *ListClustersStatusListList      `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize         int64                            `position:"Query" name:"PageSize"`
	ClusterTypeLists *ListClustersClusterTypeListList `position:"Query" type:"Repeated" name:"ClusterTypeList"`
	IsDesc           string                           `position:"Query" name:"IsDesc"`
	CreateType       string                           `position:"Query" name:"CreateType"`
	DefaultStatus    string                           `position:"Query" name:"DefaultStatus"`
	PageNumber       int64                            `position:"Query" name:"PageNumber"`
	RegionId         string                           `position:"Query" name:"RegionId"`
}

func (req *ListClustersRequest) Invoke(client goaliyun.Client) (*ListClustersResponse, error) {
	resp := &ListClustersResponse{}
	err := client.Request("emr", "ListClusters", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClustersResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Clusters   ListClustersClusterInfoList
}

type ListClustersClusterInfo struct {
	Id                  goaliyun.String
	Name                goaliyun.String
	Type                goaliyun.String
	CreateTime          goaliyun.Integer
	RunningTime         goaliyun.Integer
	Status              goaliyun.String
	ChargeType          goaliyun.String
	ExpiredTime         goaliyun.Integer
	Period              goaliyun.Integer
	HasUncompletedOrder bool
	OrderList           goaliyun.String
	CreateResource      goaliyun.String
	OrderTaskInfo       ListClustersOrderTaskInfo
	FailReason          ListClustersFailReason
}

type ListClustersOrderTaskInfo struct {
	TargetCount  goaliyun.Integer
	CurrentCount goaliyun.Integer
	OrderIdList  goaliyun.String
}

type ListClustersFailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}

type ListClustersStatusListList []string

func (list *ListClustersStatusListList) UnmarshalJSON(data []byte) error {
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

type ListClustersClusterTypeListList []string

func (list *ListClustersClusterTypeListList) UnmarshalJSON(data []byte) error {
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

type ListClustersClusterInfoList []ListClustersClusterInfo

func (list *ListClustersClusterInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClustersClusterInfo)
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
