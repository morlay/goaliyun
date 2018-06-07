package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterScriptRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterScriptRequest) Invoke(client goaliyun.Client) (*DescribeClusterScriptResponse, error) {
	resp := &DescribeClusterScriptResponse{}
	err := client.Request("emr", "DescribeClusterScript", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterScriptResponse struct {
	RequestId           goaliyun.String
	ScriptNodeInstances DescribeClusterScriptScriptNodeInstanceList
}

type DescribeClusterScriptScriptNodeInstance struct {
	NodeId    goaliyun.String
	NodeIp    goaliyun.String
	StartTime goaliyun.Integer
	EndTime   goaliyun.Integer
	Status    goaliyun.String
}

type DescribeClusterScriptScriptNodeInstanceList []DescribeClusterScriptScriptNodeInstance

func (list *DescribeClusterScriptScriptNodeInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterScriptScriptNodeInstance)
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
