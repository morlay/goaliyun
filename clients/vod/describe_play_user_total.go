package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePlayUserTotalRequest struct {
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
	OwnerId   int64  `position:"Query" name:"OwnerId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DescribePlayUserTotalRequest) Invoke(client goaliyun.Client) (*DescribePlayUserTotalResponse, error) {
	resp := &DescribePlayUserTotalResponse{}
	err := client.Request("vod", "DescribePlayUserTotal", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePlayUserTotalResponse struct {
	RequestId            goaliyun.String
	UserPlayStatisTotals DescribePlayUserTotalUserPlayStatisTotalList
}

type DescribePlayUserTotalUserPlayStatisTotal struct {
	Date         goaliyun.String
	PlayDuration goaliyun.String
	PlayRange    goaliyun.String
	VV           DescribePlayUserTotalVV
	UV           DescribePlayUserTotalUV
}

type DescribePlayUserTotalVV struct {
	Android goaliyun.String
	IOS     goaliyun.String
	Flash   goaliyun.String
	HTML5   goaliyun.String
}

type DescribePlayUserTotalUV struct {
	Android goaliyun.String
	IOS     goaliyun.String
	Flash   goaliyun.String
	HTML5   goaliyun.String
}

type DescribePlayUserTotalUserPlayStatisTotalList []DescribePlayUserTotalUserPlayStatisTotal

func (list *DescribePlayUserTotalUserPlayStatisTotalList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePlayUserTotalUserPlayStatisTotal)
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
