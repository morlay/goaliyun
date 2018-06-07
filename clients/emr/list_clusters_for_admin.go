package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClustersForAdminRequest struct {
	ResourceOwnerId  int64                                    `position:"Query" name:"ResourceOwnerId"`
	StatusLists      *ListClustersForAdminStatusListList      `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize         int64                                    `position:"Query" name:"PageSize"`
	ClusterTypeLists *ListClustersForAdminClusterTypeListList `position:"Query" type:"Repeated" name:"ClusterTypeList"`
	IsDesc           string                                   `position:"Query" name:"IsDesc"`
	UserId           string                                   `position:"Query" name:"UserId"`
	CreateType       string                                   `position:"Query" name:"CreateType"`
	DefaultStatus    string                                   `position:"Query" name:"DefaultStatus"`
	PageNumber       int64                                    `position:"Query" name:"PageNumber"`
	RegionId         string                                   `position:"Query" name:"RegionId"`
}

func (req *ListClustersForAdminRequest) Invoke(client goaliyun.Client) (*ListClustersForAdminResponse, error) {
	resp := &ListClustersForAdminResponse{}
	err := client.Request("emr", "ListClustersForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClustersForAdminResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Clusters   ListClustersForAdminClusterInfoList
}

type ListClustersForAdminClusterInfo struct {
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
	OrderTaskInfo       ListClustersForAdminOrderTaskInfo
	FailReason          ListClustersForAdminFailReason
}

type ListClustersForAdminOrderTaskInfo struct {
	TargetCount  goaliyun.Integer
	CurrentCount goaliyun.Integer
	OrderIdList  goaliyun.String
}

type ListClustersForAdminFailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}

type ListClustersForAdminStatusListList []string

func (list *ListClustersForAdminStatusListList) UnmarshalJSON(data []byte) error {
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

type ListClustersForAdminClusterTypeListList []string

func (list *ListClustersForAdminClusterTypeListList) UnmarshalJSON(data []byte) error {
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

type ListClustersForAdminClusterInfoList []ListClustersForAdminClusterInfo

func (list *ListClustersForAdminClusterInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClustersForAdminClusterInfo)
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
