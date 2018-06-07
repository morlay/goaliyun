package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFlowClusterRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListFlowClusterRequest) Invoke(client goaliyun.Client) (*ListFlowClusterResponse, error) {
	resp := &ListFlowClusterResponse{}
	err := client.Request("emr", "ListFlowCluster", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFlowClusterResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Clusters   ListFlowClusterClusterInfoList
}

type ListFlowClusterClusterInfo struct {
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
	OrderTaskInfo       ListFlowClusterOrderTaskInfo
	FailReason          ListFlowClusterFailReason
}

type ListFlowClusterOrderTaskInfo struct {
	TargetCount  goaliyun.Integer
	CurrentCount goaliyun.Integer
	OrderIdList  goaliyun.String
}

type ListFlowClusterFailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}

type ListFlowClusterClusterInfoList []ListFlowClusterClusterInfo

func (list *ListFlowClusterClusterInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowClusterClusterInfo)
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
