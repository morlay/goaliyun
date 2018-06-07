package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeFlowProjectClusterSettingRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeFlowProjectClusterSettingRequest) Invoke(client goaliyun.Client) (*DescribeFlowProjectClusterSettingResponse, error) {
	resp := &DescribeFlowProjectClusterSettingResponse{}
	err := client.Request("emr", "DescribeFlowProjectClusterSetting", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFlowProjectClusterSettingResponse struct {
	RequestId    goaliyun.String
	GmtCreate    goaliyun.Integer
	GmtModified  goaliyun.Integer
	ProjectId    goaliyun.String
	ClusterId    goaliyun.String
	DefaultUser  goaliyun.String
	DefaultQueue goaliyun.String
	UserList     DescribeFlowProjectClusterSettingUserListList
	QueueList    DescribeFlowProjectClusterSettingQueueListList
	HostList     DescribeFlowProjectClusterSettingHostListList
}

type DescribeFlowProjectClusterSettingUserListList []goaliyun.String

func (list *DescribeFlowProjectClusterSettingUserListList) UnmarshalJSON(data []byte) error {
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

type DescribeFlowProjectClusterSettingQueueListList []goaliyun.String

func (list *DescribeFlowProjectClusterSettingQueueListList) UnmarshalJSON(data []byte) error {
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

type DescribeFlowProjectClusterSettingHostListList []goaliyun.String

func (list *DescribeFlowProjectClusterSettingHostListList) UnmarshalJSON(data []byte) error {
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
