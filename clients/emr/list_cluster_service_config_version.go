package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterServiceConfigVersionRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterServiceConfigVersionRequest) Invoke(client goaliyun.Client) (*ListClusterServiceConfigVersionResponse, error) {
	resp := &ListClusterServiceConfigVersionResponse{}
	err := client.Request("emr", "ListClusterServiceConfigVersion", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterServiceConfigVersionResponse struct {
	RequestId         goaliyun.String
	PageNumber        goaliyun.Integer
	PageSize          goaliyun.Integer
	TotalCount        goaliyun.Integer
	ConfigVersionList ListClusterServiceConfigVersionConfigVersionList
}

type ListClusterServiceConfigVersionConfigVersion struct {
	ConfigVersion goaliyun.String
	Applied       bool
	CreateTime    goaliyun.Integer
	Author        goaliyun.String
	Comment       goaliyun.String
}

type ListClusterServiceConfigVersionConfigVersionList []ListClusterServiceConfigVersionConfigVersion

func (list *ListClusterServiceConfigVersionConfigVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigVersionConfigVersion)
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
