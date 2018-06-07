package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSQLLogReportsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSQLLogReportsRequest) Invoke(client goaliyun.Client) (*DescribeSQLLogReportsResponse, error) {
	resp := &DescribeSQLLogReportsResponse{}
	err := client.Request("rds", "DescribeSQLLogReports", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSQLLogReportsResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeSQLLogReportsItemList
}

type DescribeSQLLogReportsItem struct {
	ReportTime       goaliyun.String
	LatencyTopNItems DescribeSQLLogReportsLatencyTopNItemList
	QPSTopNItems     DescribeSQLLogReportsQPSTopNItemList
}

type DescribeSQLLogReportsLatencyTopNItem struct {
	SQLText         goaliyun.String
	AvgLatency      goaliyun.Integer
	SQLExecuteTimes goaliyun.Integer
}

type DescribeSQLLogReportsQPSTopNItem struct {
	SQLText         goaliyun.String
	SQLExecuteTimes goaliyun.Integer
}

type DescribeSQLLogReportsItemList []DescribeSQLLogReportsItem

func (list *DescribeSQLLogReportsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportsItem)
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

type DescribeSQLLogReportsLatencyTopNItemList []DescribeSQLLogReportsLatencyTopNItem

func (list *DescribeSQLLogReportsLatencyTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportsLatencyTopNItem)
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

type DescribeSQLLogReportsQPSTopNItemList []DescribeSQLLogReportsQPSTopNItem

func (list *DescribeSQLLogReportsQPSTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportsQPSTopNItem)
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
