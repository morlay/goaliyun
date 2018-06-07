package httpdns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetResolveStatisticsRequest struct {
	DomainName   string `position:"Query" name:"DomainName"`
	TimeSpan     int64  `position:"Query" name:"TimeSpan"`
	ProtocolName string `position:"Query" name:"ProtocolName"`
	Granularity  string `position:"Query" name:"Granularity"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *GetResolveStatisticsRequest) Invoke(client goaliyun.Client) (*GetResolveStatisticsResponse, error) {
	resp := &GetResolveStatisticsResponse{}
	err := client.Request("httpdns", "GetResolveStatistics", "2016-02-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetResolveStatisticsResponse struct {
	RequestId  goaliyun.String
	DataPoints GetResolveStatisticsDataPointList
}

type GetResolveStatisticsDataPoint struct {
	Time  goaliyun.Integer
	Count goaliyun.Integer
}

type GetResolveStatisticsDataPointList []GetResolveStatisticsDataPoint

func (list *GetResolveStatisticsDataPointList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResolveStatisticsDataPoint)
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
