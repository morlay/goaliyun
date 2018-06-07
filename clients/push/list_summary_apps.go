package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListSummaryAppsRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListSummaryAppsRequest) Invoke(client goaliyun.Client) (*ListSummaryAppsResponse, error) {
	resp := &ListSummaryAppsResponse{}
	err := client.Request("push", "ListSummaryApps", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListSummaryAppsResponse struct {
	RequestId       goaliyun.String
	SummaryAppInfos ListSummaryAppsSummaryAppInfoList
}

type ListSummaryAppsSummaryAppInfo struct {
	AppName goaliyun.String
	AppKey  goaliyun.Integer
}

type ListSummaryAppsSummaryAppInfoList []ListSummaryAppsSummaryAppInfo

func (list *ListSummaryAppsSummaryAppInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSummaryAppsSummaryAppInfo)
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
