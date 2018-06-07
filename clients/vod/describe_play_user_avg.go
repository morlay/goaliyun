package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePlayUserAvgRequest struct {
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
	OwnerId   int64  `position:"Query" name:"OwnerId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DescribePlayUserAvgRequest) Invoke(client goaliyun.Client) (*DescribePlayUserAvgResponse, error) {
	resp := &DescribePlayUserAvgResponse{}
	err := client.Request("vod", "DescribePlayUserAvg", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePlayUserAvgResponse struct {
	RequestId          goaliyun.String
	UserPlayStatisAvgs DescribePlayUserAvgUserPlayStatisAvgList
}

type DescribePlayUserAvgUserPlayStatisAvg struct {
	Date            goaliyun.String
	AvgPlayDuration goaliyun.String
	AvgPlayCount    goaliyun.String
}

type DescribePlayUserAvgUserPlayStatisAvgList []DescribePlayUserAvgUserPlayStatisAvg

func (list *DescribePlayUserAvgUserPlayStatisAvgList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePlayUserAvgUserPlayStatisAvg)
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
