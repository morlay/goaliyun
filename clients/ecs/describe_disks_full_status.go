package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDisksFullStatusRequest struct {
	EventIds             *DescribeDisksFullStatusEventIdList `position:"Query" type:"Repeated" name:"EventId"`
	ResourceOwnerId      int64                               `position:"Query" name:"ResourceOwnerId"`
	PageNumber           int64                               `position:"Query" name:"PageNumber"`
	EventTimeStart       string                              `position:"Query" name:"EventTimeStart"`
	PageSize             int64                               `position:"Query" name:"PageSize"`
	DiskIds              *DescribeDisksFullStatusDiskIdList  `position:"Query" type:"Repeated" name:"DiskId"`
	ResourceOwnerAccount string                              `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                              `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                               `position:"Query" name:"OwnerId"`
	EventTimeEnd         string                              `position:"Query" name:"EventTimeEnd"`
	HealthStatus         string                              `position:"Query" name:"HealthStatus"`
	EventType            string                              `position:"Query" name:"EventType"`
	Status               string                              `position:"Query" name:"Status"`
	RegionId             string                              `position:"Query" name:"RegionId"`
}

func (req *DescribeDisksFullStatusRequest) Invoke(client goaliyun.Client) (*DescribeDisksFullStatusResponse, error) {
	resp := &DescribeDisksFullStatusResponse{}
	err := client.Request("ecs", "DescribeDisksFullStatus", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDisksFullStatusResponse struct {
	RequestId         goaliyun.String
	TotalCount        goaliyun.Integer
	PageNumber        goaliyun.Integer
	PageSize          goaliyun.Integer
	DiskFullStatusSet DescribeDisksFullStatusDiskFullStatusTypeList
}

type DescribeDisksFullStatusDiskFullStatusType struct {
	DiskId       goaliyun.String
	DiskEventSet DescribeDisksFullStatusDiskEventTypeList
	Status       DescribeDisksFullStatusStatus
	HealthStatus DescribeDisksFullStatusHealthStatus
}

type DescribeDisksFullStatusDiskEventType struct {
	EventId      goaliyun.String
	EventTime    goaliyun.String
	EventEndTime goaliyun.String
	EventType    DescribeDisksFullStatusEventType
}

type DescribeDisksFullStatusEventType struct {
	Code goaliyun.Integer
	Name goaliyun.String
}

type DescribeDisksFullStatusStatus struct {
	Code goaliyun.Integer
	Name goaliyun.String
}

type DescribeDisksFullStatusHealthStatus struct {
	Code goaliyun.Integer
	Name goaliyun.String
}

type DescribeDisksFullStatusEventIdList []string

func (list *DescribeDisksFullStatusEventIdList) UnmarshalJSON(data []byte) error {
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

type DescribeDisksFullStatusDiskIdList []string

func (list *DescribeDisksFullStatusDiskIdList) UnmarshalJSON(data []byte) error {
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

type DescribeDisksFullStatusDiskFullStatusTypeList []DescribeDisksFullStatusDiskFullStatusType

func (list *DescribeDisksFullStatusDiskFullStatusTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksFullStatusDiskFullStatusType)
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

type DescribeDisksFullStatusDiskEventTypeList []DescribeDisksFullStatusDiskEventType

func (list *DescribeDisksFullStatusDiskEventTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksFullStatusDiskEventType)
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
