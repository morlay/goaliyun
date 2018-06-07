package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type NodeListRequest struct {
	HostName         string `position:"Query" name:"HostName"`
	InstanceIds      string `position:"Query" name:"InstanceIds"`
	InstanceRegionId string `position:"Query" name:"InstanceRegionId"`
	PageSize         int64  `position:"Query" name:"PageSize"`
	KeyWord          string `position:"Query" name:"KeyWord"`
	UserId           int64  `position:"Query" name:"UserId"`
	PageNumber       int64  `position:"Query" name:"PageNumber"`
	SerialNumbers    string `position:"Query" name:"SerialNumbers"`
	Status           string `position:"Query" name:"Status"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *NodeListRequest) Invoke(client goaliyun.Client) (*NodeListResponse, error) {
	resp := &NodeListResponse{}
	err := client.Request("cms", "NodeList", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type NodeListResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	PageTotal    goaliyun.Integer
	Total        goaliyun.Integer
	Nodes        NodeListNodeList
}

type NodeListNode struct {
	InstanceId       goaliyun.String
	SerialNumber     goaliyun.String
	HostName         goaliyun.String
	AliUid           goaliyun.Integer
	OperatingSystem  goaliyun.String
	IpGroup          goaliyun.String
	Region           goaliyun.String
	TianjimonVersion goaliyun.String
	EipAddress       goaliyun.String
	EipId            goaliyun.String
	AliyunHost       bool
	NatIp            goaliyun.String
	NetworkType      goaliyun.String
}

type NodeListNodeList []NodeListNode

func (list *NodeListNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]NodeListNode)
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
