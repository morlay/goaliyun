package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SearchLogRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostInnerIp     string `position:"Query" name:"HostInnerIp"`
	LogstoreName    string `position:"Query" name:"LogstoreName"`
	FromTimestamp   int64  `position:"Query" name:"FromTimestamp"`
	Offset          int64  `position:"Query" name:"Offset"`
	Line            int64  `position:"Query" name:"Line"`
	ToTimestamp     int64  `position:"Query" name:"ToTimestamp"`
	SlsQueryString  string `position:"Query" name:"SlsQueryString"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	Reverse         string `position:"Query" name:"Reverse"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *SearchLogRequest) Invoke(client goaliyun.Client) (*SearchLogResponse, error) {
	resp := &SearchLogResponse{}
	err := client.Request("emr", "SearchLog", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SearchLogResponse struct {
	RequestId      goaliyun.String
	Completed      bool
	SlsLogItemList SearchLogSlsLogItemList
}

type SearchLogSlsLogItem struct {
	Timestamp goaliyun.Integer
	SourceIp  goaliyun.String
	HostName  goaliyun.String
	Path      goaliyun.String
	Content   goaliyun.String
}

type SearchLogSlsLogItemList []SearchLogSlsLogItem

func (list *SearchLogSlsLogItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchLogSlsLogItem)
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
