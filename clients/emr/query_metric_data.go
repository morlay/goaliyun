package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryMetricDataRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Period          int64  `position:"Query" name:"Period"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	StartTimeStamp  int64  `position:"Query" name:"StartTimeStamp"`
	MetricName      string `position:"Query" name:"MetricName"`
	HostRole        string `position:"Query" name:"HostRole"`
	EndTimeStamp    int64  `position:"Query" name:"EndTimeStamp"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryMetricDataRequest) Invoke(client goaliyun.Client) (*QueryMetricDataResponse, error) {
	resp := &QueryMetricDataResponse{}
	err := client.Request("emr", "QueryMetricData", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMetricDataResponse struct {
	RequestId  goaliyun.String
	Datapoints QueryMetricDataCmsDataPointList
}

type QueryMetricDataCmsDataPoint struct {
	Role      goaliyun.String
	Maximum   goaliyun.Float
	Minimum   goaliyun.Float
	Average   goaliyun.Float
	Timestamp goaliyun.Integer
}

type QueryMetricDataCmsDataPointList []QueryMetricDataCmsDataPoint

func (list *QueryMetricDataCmsDataPointList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMetricDataCmsDataPoint)
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
