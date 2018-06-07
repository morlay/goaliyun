package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateFlowProjectClusterSettingRequest struct {
	UserLists       *CreateFlowProjectClusterSettingUserListList  `position:"Query" type:"Repeated" name:"UserList"`
	ResourceOwnerId int64                                         `position:"Query" name:"ResourceOwnerId"`
	QueueLists      *CreateFlowProjectClusterSettingQueueListList `position:"Query" type:"Repeated" name:"QueueList"`
	HostLists       *CreateFlowProjectClusterSettingHostListList  `position:"Query" type:"Repeated" name:"HostList"`
	ClusterId       string                                        `position:"Query" name:"ClusterId"`
	DefaultQueue    string                                        `position:"Query" name:"DefaultQueue"`
	ProjectId       string                                        `position:"Query" name:"ProjectId"`
	DefaultUser     string                                        `position:"Query" name:"DefaultUser"`
	RegionId        string                                        `position:"Query" name:"RegionId"`
}

func (req *CreateFlowProjectClusterSettingRequest) Invoke(client goaliyun.Client) (*CreateFlowProjectClusterSettingResponse, error) {
	resp := &CreateFlowProjectClusterSettingResponse{}
	err := client.Request("emr", "CreateFlowProjectClusterSetting", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateFlowProjectClusterSettingResponse struct {
	RequestId goaliyun.String
	Data      bool
}

type CreateFlowProjectClusterSettingUserListList []string

func (list *CreateFlowProjectClusterSettingUserListList) UnmarshalJSON(data []byte) error {
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

type CreateFlowProjectClusterSettingQueueListList []string

func (list *CreateFlowProjectClusterSettingQueueListList) UnmarshalJSON(data []byte) error {
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

type CreateFlowProjectClusterSettingHostListList []string

func (list *CreateFlowProjectClusterSettingHostListList) UnmarshalJSON(data []byte) error {
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
