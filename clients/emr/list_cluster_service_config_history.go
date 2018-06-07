package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterServiceConfigHistoryRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	ConfigVersion   string `position:"Query" name:"ConfigVersion"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterServiceConfigHistoryRequest) Invoke(client goaliyun.Client) (*ListClusterServiceConfigHistoryResponse, error) {
	resp := &ListClusterServiceConfigHistoryResponse{}
	err := client.Request("emr", "ListClusterServiceConfigHistory", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterServiceConfigHistoryResponse struct {
	RequestId         goaliyun.String
	TotalCount        goaliyun.Integer
	PageNumber        goaliyun.Integer
	PageSize          goaliyun.Integer
	ConfigHistoryList ListClusterServiceConfigHistoryConfigHistoryList
}

type ListClusterServiceConfigHistoryConfigHistory struct {
	ServiceName    goaliyun.String
	ConfigVersion  goaliyun.String
	ConfigFileName goaliyun.String
	ConfigItemName goaliyun.String
	NewValue       goaliyun.String
	OldValue       goaliyun.String
	Applied        bool
	CreateTime     goaliyun.Integer
	Author         goaliyun.String
	Comment        goaliyun.String
}

type ListClusterServiceConfigHistoryConfigHistoryList []ListClusterServiceConfigHistoryConfigHistory

func (list *ListClusterServiceConfigHistoryConfigHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigHistoryConfigHistory)
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
