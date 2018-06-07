package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyFlowProjectClusterSettingRequest struct {
	UserLists       *ModifyFlowProjectClusterSettingUserListList  `position:"Query" type:"Repeated" name:"UserList"`
	ResourceOwnerId int64                                         `position:"Query" name:"ResourceOwnerId"`
	QueueLists      *ModifyFlowProjectClusterSettingQueueListList `position:"Query" type:"Repeated" name:"QueueList"`
	HostLists       *ModifyFlowProjectClusterSettingHostListList  `position:"Query" type:"Repeated" name:"HostList"`
	ClusterId       string                                        `position:"Query" name:"ClusterId"`
	DefaultQueue    string                                        `position:"Query" name:"DefaultQueue"`
	ProjectId       string                                        `position:"Query" name:"ProjectId"`
	DefaultUser     string                                        `position:"Query" name:"DefaultUser"`
	RegionId        string                                        `position:"Query" name:"RegionId"`
}

func (req *ModifyFlowProjectClusterSettingRequest) Invoke(client goaliyun.Client) (*ModifyFlowProjectClusterSettingResponse, error) {
	resp := &ModifyFlowProjectClusterSettingResponse{}
	err := client.Request("emr", "ModifyFlowProjectClusterSetting", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyFlowProjectClusterSettingResponse struct {
	RequestId goaliyun.String
	Data      bool
}

type ModifyFlowProjectClusterSettingUserListList []string

func (list *ModifyFlowProjectClusterSettingUserListList) UnmarshalJSON(data []byte) error {
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

type ModifyFlowProjectClusterSettingQueueListList []string

func (list *ModifyFlowProjectClusterSettingQueueListList) UnmarshalJSON(data []byte) error {
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

type ModifyFlowProjectClusterSettingHostListList []string

func (list *ModifyFlowProjectClusterSettingHostListList) UnmarshalJSON(data []byte) error {
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
