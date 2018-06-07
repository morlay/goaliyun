package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFlowProjectClusterSettingRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListFlowProjectClusterSettingRequest) Invoke(client goaliyun.Client) (*ListFlowProjectClusterSettingResponse, error) {
	resp := &ListFlowProjectClusterSettingResponse{}
	err := client.Request("emr", "ListFlowProjectClusterSetting", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFlowProjectClusterSettingResponse struct {
	RequestId       goaliyun.String
	PageNumber      goaliyun.Integer
	PageSize        goaliyun.Integer
	Total           goaliyun.Integer
	ClusterSettings ListFlowProjectClusterSettingClusterSettingList
}

type ListFlowProjectClusterSettingClusterSetting struct {
	GmtCreate    goaliyun.Integer
	GmtModified  goaliyun.Integer
	ProjectId    goaliyun.String
	ClusterId    goaliyun.String
	DefaultUser  goaliyun.String
	DefaultQueue goaliyun.String
	UserList     ListFlowProjectClusterSettingUserListList
	QueueList    ListFlowProjectClusterSettingQueueListList
	HostList     ListFlowProjectClusterSettingHostListList
}

type ListFlowProjectClusterSettingClusterSettingList []ListFlowProjectClusterSettingClusterSetting

func (list *ListFlowProjectClusterSettingClusterSettingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowProjectClusterSettingClusterSetting)
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

type ListFlowProjectClusterSettingUserListList []goaliyun.String

func (list *ListFlowProjectClusterSettingUserListList) UnmarshalJSON(data []byte) error {
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

type ListFlowProjectClusterSettingQueueListList []goaliyun.String

func (list *ListFlowProjectClusterSettingQueueListList) UnmarshalJSON(data []byte) error {
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

type ListFlowProjectClusterSettingHostListList []goaliyun.String

func (list *ListFlowProjectClusterSettingHostListList) UnmarshalJSON(data []byte) error {
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
