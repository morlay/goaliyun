package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeEventsRequest struct {
	EventId              string `position:"Query" name:"EventId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	PlanTime             string `position:"Query" name:"PlanTime"`
	ExpireTime           string `position:"Query" name:"ExpireTime"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	EventType            string `position:"Query" name:"EventType"`
	Status               string `position:"Query" name:"Status"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeEventsRequest) Invoke(client goaliyun.Client) (*DescribeEventsResponse, error) {
	resp := &DescribeEventsResponse{}
	err := client.Request("ecs", "DescribeEvents", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeEventsResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	Events     DescribeEventsEventList
}

type DescribeEventsEvent struct {
	ResourceId    goaliyun.String
	EventType     goaliyun.String
	EventCategory goaliyun.String
	Status        goaliyun.String
	SupportModify goaliyun.String
	PlanTime      goaliyun.String
	ExpireTime    goaliyun.String
	EventId       goaliyun.String
}

type DescribeEventsEventList []DescribeEventsEvent

func (list *DescribeEventsEventList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEventsEvent)
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
