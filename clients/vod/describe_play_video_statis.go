package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePlayVideoStatisRequest struct {
	EndTime   string `position:"Query" name:"EndTime"`
	VideoId   string `position:"Query" name:"VideoId"`
	StartTime string `position:"Query" name:"StartTime"`
	OwnerId   int64  `position:"Query" name:"OwnerId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DescribePlayVideoStatisRequest) Invoke(client goaliyun.Client) (*DescribePlayVideoStatisResponse, error) {
	resp := &DescribePlayVideoStatisResponse{}
	err := client.Request("vod", "DescribePlayVideoStatis", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePlayVideoStatisResponse struct {
	RequestId              goaliyun.String
	VideoPlayStatisDetails DescribePlayVideoStatisVideoPlayStatisDetailList
}

type DescribePlayVideoStatisVideoPlayStatisDetail struct {
	Date         goaliyun.String
	PlayDuration goaliyun.String
	VV           goaliyun.String
	UV           goaliyun.String
	PlayRange    goaliyun.String
	Title        goaliyun.String
}

type DescribePlayVideoStatisVideoPlayStatisDetailList []DescribePlayVideoStatisVideoPlayStatisDetail

func (list *DescribePlayVideoStatisVideoPlayStatisDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePlayVideoStatisVideoPlayStatisDetail)
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
