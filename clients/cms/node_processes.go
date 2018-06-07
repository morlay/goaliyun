package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type NodeProcessesRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *NodeProcessesRequest) Invoke(client goaliyun.Client) (*NodeProcessesResponse, error) {
	resp := &NodeProcessesResponse{}
	err := client.Request("cms", "NodeProcesses", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type NodeProcessesResponse struct {
	ErrorCode     goaliyun.Integer
	ErrorMessage  goaliyun.String
	Success       bool
	RequestId     goaliyun.String
	NodeProcesses NodeProcessesNodeProcessList
}

type NodeProcessesNodeProcess struct {
	Id          goaliyun.Integer
	Name        goaliyun.String
	InstanceId  goaliyun.String
	ProcessName goaliyun.String
	ProcessUser goaliyun.String
	Command     goaliyun.String
}

type NodeProcessesNodeProcessList []NodeProcessesNodeProcess

func (list *NodeProcessesNodeProcessList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]NodeProcessesNodeProcess)
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
