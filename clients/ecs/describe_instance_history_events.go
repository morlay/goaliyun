package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstanceHistoryEventsRequest struct {
	EventIds                  *DescribeInstanceHistoryEventsEventIdList                  `position:"Query" type:"Repeated" name:"EventId"`
	ResourceOwnerId           int64                                                      `position:"Query" name:"ResourceOwnerId"`
	EventCycleStatus          string                                                     `position:"Query" name:"EventCycleStatus"`
	PageNumber                int64                                                      `position:"Query" name:"PageNumber"`
	PageSize                  int64                                                      `position:"Query" name:"PageSize"`
	InstanceEventCycleStatuss *DescribeInstanceHistoryEventsInstanceEventCycleStatusList `position:"Query" type:"Repeated" name:"InstanceEventCycleStatus"`
	EventPublishTimeEnd       string                                                     `position:"Query" name:"EventPublishTimeEnd"`
	InstanceEventTypes        *DescribeInstanceHistoryEventsInstanceEventTypeList        `position:"Query" type:"Repeated" name:"InstanceEventType"`
	ResourceOwnerAccount      string                                                     `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string                                                     `position:"Query" name:"OwnerAccount"`
	NotBeforeStart            string                                                     `position:"Query" name:"NotBeforeStart"`
	OwnerId                   int64                                                      `position:"Query" name:"OwnerId"`
	EventPublishTimeStart     string                                                     `position:"Query" name:"EventPublishTimeStart"`
	InstanceId                string                                                     `position:"Query" name:"InstanceId"`
	NotBeforeEnd              string                                                     `position:"Query" name:"NotBeforeEnd"`
	EventType                 string                                                     `position:"Query" name:"EventType"`
	RegionId                  string                                                     `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceHistoryEventsRequest) Invoke(client goaliyun.Client) (*DescribeInstanceHistoryEventsResponse, error) {
	resp := &DescribeInstanceHistoryEventsResponse{}
	err := client.Request("ecs", "DescribeInstanceHistoryEvents", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceHistoryEventsResponse struct {
	RequestId              goaliyun.String
	TotalCount             goaliyun.Integer
	PageNumber             goaliyun.Integer
	PageSize               goaliyun.Integer
	InstanceSystemEventSet DescribeInstanceHistoryEventsInstanceSystemEventTypeList
}

type DescribeInstanceHistoryEventsInstanceSystemEventType struct {
	InstanceId       goaliyun.String
	EventId          goaliyun.String
	EventPublishTime goaliyun.String
	NotBefore        goaliyun.String
	EventFinishTime  goaliyun.String
	EventType        DescribeInstanceHistoryEventsEventType
	EventCycleStatus DescribeInstanceHistoryEventsEventCycleStatus
}

type DescribeInstanceHistoryEventsEventType struct {
	Code goaliyun.Integer
	Name goaliyun.String
}

type DescribeInstanceHistoryEventsEventCycleStatus struct {
	Code goaliyun.Integer
	Name goaliyun.String
}

type DescribeInstanceHistoryEventsEventIdList []string

func (list *DescribeInstanceHistoryEventsEventIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeInstanceHistoryEventsInstanceEventCycleStatusList []string

func (list *DescribeInstanceHistoryEventsInstanceEventCycleStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeInstanceHistoryEventsInstanceEventTypeList []string

func (list *DescribeInstanceHistoryEventsInstanceEventTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeInstanceHistoryEventsInstanceSystemEventTypeList []DescribeInstanceHistoryEventsInstanceSystemEventType

func (list *DescribeInstanceHistoryEventsInstanceSystemEventTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceHistoryEventsInstanceSystemEventType)
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
