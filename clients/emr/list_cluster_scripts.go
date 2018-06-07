package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterScriptsRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterScriptsRequest) Invoke(client goaliyun.Client) (*ListClusterScriptsResponse, error) {
	resp := &ListClusterScriptsResponse{}
	err := client.Request("emr", "ListClusterScripts", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterScriptsResponse struct {
	RequestId      goaliyun.String
	ClusterScripts ListClusterScriptsClusterScriptList
}

type ListClusterScriptsClusterScript struct {
	Id        goaliyun.String
	Name      goaliyun.String
	StartTime goaliyun.Integer
	EndTime   goaliyun.Integer
	Path      goaliyun.String
	Args      goaliyun.String
	Status    goaliyun.String
}

type ListClusterScriptsClusterScriptList []ListClusterScriptsClusterScript

func (list *ListClusterScriptsClusterScriptList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterScriptsClusterScript)
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
