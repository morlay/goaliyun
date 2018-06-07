package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePlayTopVideosRequest struct {
	BizDate  string `position:"Query" name:"BizDate"`
	PageNo   int64  `position:"Query" name:"PageNo"`
	PageSize int64  `position:"Query" name:"PageSize"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribePlayTopVideosRequest) Invoke(client goaliyun.Client) (*DescribePlayTopVideosResponse, error) {
	resp := &DescribePlayTopVideosResponse{}
	err := client.Request("vod", "DescribePlayTopVideos", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePlayTopVideosResponse struct {
	RequestId     goaliyun.String
	PageNo        goaliyun.Integer
	PageSize      goaliyun.Integer
	TotalNum      goaliyun.Integer
	TopPlayVideos DescribePlayTopVideosTopPlayVideoStatisList
}

type DescribePlayTopVideosTopPlayVideoStatis struct {
	PlayDuration goaliyun.String
	VV           goaliyun.String
	UV           goaliyun.String
	VideoId      goaliyun.String
	Title        goaliyun.String
}

type DescribePlayTopVideosTopPlayVideoStatisList []DescribePlayTopVideosTopPlayVideoStatis

func (list *DescribePlayTopVideosTopPlayVideoStatisList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePlayTopVideosTopPlayVideoStatis)
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
