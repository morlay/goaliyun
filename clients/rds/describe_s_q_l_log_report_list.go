package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSQLLogReportListRequest struct {
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

func (req *DescribeSQLLogReportListRequest) Invoke(client goaliyun.Client) (*DescribeSQLLogReportListResponse, error) {
	resp := &DescribeSQLLogReportListResponse{}
	err := client.Request("rds", "DescribeSQLLogReportList", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSQLLogReportListResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeSQLLogReportListItemList
}

type DescribeSQLLogReportListItem struct {
	ReportTime       goaliyun.String
	LatencyTopNItems DescribeSQLLogReportListLatencyTopNItemList
	QPSTopNItems     DescribeSQLLogReportListQPSTopNItemList
}

type DescribeSQLLogReportListLatencyTopNItem struct {
	SQLText         goaliyun.String
	AvgLatency      goaliyun.Integer
	SQLExecuteTimes goaliyun.Integer
}

type DescribeSQLLogReportListQPSTopNItem struct {
	SQLText         goaliyun.String
	SQLExecuteTimes goaliyun.Integer
}

type DescribeSQLLogReportListItemList []DescribeSQLLogReportListItem

func (list *DescribeSQLLogReportListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportListItem)
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

type DescribeSQLLogReportListLatencyTopNItemList []DescribeSQLLogReportListLatencyTopNItem

func (list *DescribeSQLLogReportListLatencyTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportListLatencyTopNItem)
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

type DescribeSQLLogReportListQPSTopNItemList []DescribeSQLLogReportListQPSTopNItem

func (list *DescribeSQLLogReportListQPSTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportListQPSTopNItem)
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
