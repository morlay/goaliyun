package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterServiceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterServiceRequest) Invoke(client goaliyun.Client) (*ListClusterServiceResponse, error) {
	resp := &ListClusterServiceResponse{}
	err := client.Request("emr", "ListClusterService", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterServiceResponse struct {
	RequestId          goaliyun.String
	TotalCount         goaliyun.Integer
	PageNumber         goaliyun.Integer
	PageSize           goaliyun.Integer
	ClusterServiceList ListClusterServiceClusterServiceList
}

type ListClusterServiceClusterService struct {
	ServiceName        goaliyun.String
	ServiceDisplayName goaliyun.String
	ServiceVersion     goaliyun.String
	InstallStatus      bool
	ClientType         bool
	ServiceStatus      goaliyun.String
	HealthStatus       goaliyun.String
	NeedRestartInfo    goaliyun.String
	NotStartInfo       goaliyun.String
	AbnormalNum        goaliyun.Integer
	StoppedNum         goaliyun.Integer
	NeedRestartNum     goaliyun.Integer
	ServiceActionList  ListClusterServiceServiceActionList
}

type ListClusterServiceServiceAction struct {
	ServiceName   goaliyun.String
	ComponentName goaliyun.String
	ActionName    goaliyun.String
	Command       goaliyun.String
	DisplayName   goaliyun.String
}

type ListClusterServiceClusterServiceList []ListClusterServiceClusterService

func (list *ListClusterServiceClusterServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceClusterService)
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

type ListClusterServiceServiceActionList []ListClusterServiceServiceAction

func (list *ListClusterServiceServiceActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceServiceAction)
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
