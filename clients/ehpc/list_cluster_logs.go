package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterLogsRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	ClusterId  string `position:"Query" name:"ClusterId"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListClusterLogsRequest) Invoke(client goaliyun.Client) (*ListClusterLogsResponse, error) {
	resp := &ListClusterLogsResponse{}
	err := client.Request("ehpc", "ListClusterLogs", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterLogsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	ClusterId  goaliyun.String
	Logs       ListClusterLogsLogInfoList
}

type ListClusterLogsLogInfo struct {
	Operation  goaliyun.String
	Level      goaliyun.String
	Message    goaliyun.String
	CreateTime goaliyun.String
}

type ListClusterLogsLogInfoList []ListClusterLogsLogInfo

func (list *ListClusterLogsLogInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterLogsLogInfo)
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
