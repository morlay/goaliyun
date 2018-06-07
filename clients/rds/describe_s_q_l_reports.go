package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSQLReportsRequest struct {
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

func (req *DescribeSQLReportsRequest) Invoke(client goaliyun.Client) (*DescribeSQLReportsResponse, error) {
	resp := &DescribeSQLReportsResponse{}
	err := client.Request("rds", "DescribeSQLReports", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSQLReportsResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeSQLReportsItemList
}

type DescribeSQLReportsItem struct {
	ReportTime       goaliyun.String
	LatencyTopNItems DescribeSQLReportsLatencyTopNItemList
	QPSTopNItems     DescribeSQLReportsQPSTopNItemList
}

type DescribeSQLReportsLatencyTopNItem struct {
	SQLText         goaliyun.String
	AvgLatency      goaliyun.Integer
	SQLExecuteTimes goaliyun.Integer
}

type DescribeSQLReportsQPSTopNItem struct {
	SQLText         goaliyun.String
	SQLExecuteTimes goaliyun.Integer
}

type DescribeSQLReportsItemList []DescribeSQLReportsItem

func (list *DescribeSQLReportsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLReportsItem)
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

type DescribeSQLReportsLatencyTopNItemList []DescribeSQLReportsLatencyTopNItem

func (list *DescribeSQLReportsLatencyTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLReportsLatencyTopNItem)
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

type DescribeSQLReportsQPSTopNItemList []DescribeSQLReportsQPSTopNItem

func (list *DescribeSQLReportsQPSTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLReportsQPSTopNItem)
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
