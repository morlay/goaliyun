package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSuspiciousEventsRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeSuspiciousEventsRequest) Invoke(client goaliyun.Client) (*DescribeSuspiciousEventsResponse, error) {
	resp := &DescribeSuspiciousEventsResponse{}
	err := client.Request("aegis", "DescribeSuspiciousEvents", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSuspiciousEventsResponse struct {
	RequestId        goaliyun.String
	PageSize         goaliyun.Integer
	TotalCount       goaliyun.Integer
	CurrentPage      goaliyun.Integer
	HttpStatusCode   goaliyun.Integer
	SuspiciousEvents DescribeSuspiciousEventsSuspiciousEventList
}

type DescribeSuspiciousEventsSuspiciousEventList []goaliyun.String

func (list *DescribeSuspiciousEventsSuspiciousEventList) UnmarshalJSON(data []byte) error {
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
