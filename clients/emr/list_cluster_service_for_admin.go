package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterServiceForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterServiceForAdminRequest) Invoke(client goaliyun.Client) (*ListClusterServiceForAdminResponse, error) {
	resp := &ListClusterServiceForAdminResponse{}
	err := client.Request("emr", "ListClusterServiceForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterServiceForAdminResponse struct {
	RequestId          goaliyun.String
	TotalCount         goaliyun.Integer
	PageNumber         goaliyun.Integer
	PageSize           goaliyun.Integer
	ClusterServiceList ListClusterServiceForAdminClusterServiceList
}

type ListClusterServiceForAdminClusterService struct {
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
	ServiceActionList  ListClusterServiceForAdminServiceActionList
}

type ListClusterServiceForAdminServiceAction struct {
	ServiceName   goaliyun.String
	ComponentName goaliyun.String
	ActionName    goaliyun.String
	Command       goaliyun.String
	DisplayName   goaliyun.String
}

type ListClusterServiceForAdminClusterServiceList []ListClusterServiceForAdminClusterService

func (list *ListClusterServiceForAdminClusterServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceForAdminClusterService)
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

type ListClusterServiceForAdminServiceActionList []ListClusterServiceForAdminServiceAction

func (list *ListClusterServiceForAdminServiceActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceForAdminServiceAction)
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
