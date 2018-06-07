package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstancesFullStatusRequest struct {
	EventIds              *DescribeInstancesFullStatusEventIdList           `position:"Query" type:"Repeated" name:"EventId"`
	ResourceOwnerId       int64                                             `position:"Query" name:"ResourceOwnerId"`
	PageNumber            int64                                             `position:"Query" name:"PageNumber"`
	PageSize              int64                                             `position:"Query" name:"PageSize"`
	EventPublishTimeEnd   string                                            `position:"Query" name:"EventPublishTimeEnd"`
	InstanceEventTypes    *DescribeInstancesFullStatusInstanceEventTypeList `position:"Query" type:"Repeated" name:"InstanceEventType"`
	ResourceOwnerAccount  string                                            `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string                                            `position:"Query" name:"OwnerAccount"`
	NotBeforeStart        string                                            `position:"Query" name:"NotBeforeStart"`
	OwnerId               int64                                             `position:"Query" name:"OwnerId"`
	EventPublishTimeStart string                                            `position:"Query" name:"EventPublishTimeStart"`
	InstanceIds           *DescribeInstancesFullStatusInstanceIdList        `position:"Query" type:"Repeated" name:"InstanceId"`
	NotBeforeEnd          string                                            `position:"Query" name:"NotBeforeEnd"`
	HealthStatus          string                                            `position:"Query" name:"HealthStatus"`
	EventType             string                                            `position:"Query" name:"EventType"`
	Status                string                                            `position:"Query" name:"Status"`
	RegionId              string                                            `position:"Query" name:"RegionId"`
}

func (req *DescribeInstancesFullStatusRequest) Invoke(client goaliyun.Client) (*DescribeInstancesFullStatusResponse, error) {
	resp := &DescribeInstancesFullStatusResponse{}
	err := client.Request("ecs", "DescribeInstancesFullStatus", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstancesFullStatusResponse struct {
	RequestId             goaliyun.String
	TotalCount            goaliyun.Integer
	PageNumber            goaliyun.Integer
	PageSize              goaliyun.Integer
	InstanceFullStatusSet DescribeInstancesFullStatusInstanceFullStatusTypeList
}

type DescribeInstancesFullStatusInstanceFullStatusType struct {
	InstanceId              goaliyun.String
	ScheduledSystemEventSet DescribeInstancesFullStatusScheduledSystemEventTypeList
	Status                  DescribeInstancesFullStatusStatus
	HealthStatus            DescribeInstancesFullStatusHealthStatus
}

type DescribeInstancesFullStatusScheduledSystemEventType struct {
	EventId          goaliyun.String
	EventPublishTime goaliyun.String
	NotBefore        goaliyun.String
	EventCycleStatus DescribeInstancesFullStatusEventCycleStatus
	EventType        DescribeInstancesFullStatusEventType
}

type DescribeInstancesFullStatusEventCycleStatus struct {
	Code goaliyun.Integer
	Name goaliyun.String
}

type DescribeInstancesFullStatusEventType struct {
	Code goaliyun.Integer
	Name goaliyun.String
}

type DescribeInstancesFullStatusStatus struct {
	Code goaliyun.Integer
	Name goaliyun.String
}

type DescribeInstancesFullStatusHealthStatus struct {
	Code goaliyun.Integer
	Name goaliyun.String
}

type DescribeInstancesFullStatusEventIdList []string

func (list *DescribeInstancesFullStatusEventIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesFullStatusInstanceEventTypeList []string

func (list *DescribeInstancesFullStatusInstanceEventTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesFullStatusInstanceIdList []string

func (list *DescribeInstancesFullStatusInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesFullStatusInstanceFullStatusTypeList []DescribeInstancesFullStatusInstanceFullStatusType

func (list *DescribeInstancesFullStatusInstanceFullStatusTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesFullStatusInstanceFullStatusType)
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

type DescribeInstancesFullStatusScheduledSystemEventTypeList []DescribeInstancesFullStatusScheduledSystemEventType

func (list *DescribeInstancesFullStatusScheduledSystemEventTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesFullStatusScheduledSystemEventType)
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
