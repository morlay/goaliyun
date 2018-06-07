package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type NodeStatusListRequest struct {
	InstanceIds string `position:"Query" name:"InstanceIds"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *NodeStatusListRequest) Invoke(client goaliyun.Client) (*NodeStatusListResponse, error) {
	resp := &NodeStatusListResponse{}
	err := client.Request("cms", "NodeStatusList", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type NodeStatusListResponse struct {
	ErrorCode      goaliyun.Integer
	ErrorMessage   goaliyun.String
	Success        bool
	RequestId      goaliyun.String
	NodeStatusList NodeStatusListNodeStatusList
}

type NodeStatusListNodeStatus struct {
	InstanceId  goaliyun.String
	AutoInstall bool
	Status      goaliyun.String
}

type NodeStatusListNodeStatusList []NodeStatusListNodeStatus

func (list *NodeStatusListNodeStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]NodeStatusListNodeStatus)
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
