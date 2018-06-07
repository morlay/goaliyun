package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCdnMonitorDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeCdnMonitorDataRequest) Invoke(client goaliyun.Client) (*DescribeCdnMonitorDataResponse, error) {
	resp := &DescribeCdnMonitorDataResponse{}
	err := client.Request("cdn", "DescribeCdnMonitorData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCdnMonitorDataResponse struct {
	RequestId       goaliyun.String
	DomainName      goaliyun.String
	MonitorInterval goaliyun.Integer
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	MonitorDatas    DescribeCdnMonitorDataCDNMonitorDataList
}

type DescribeCdnMonitorDataCDNMonitorData struct {
	TimeStamp         goaliyun.String
	QueryPerSecond    goaliyun.String
	BytesPerSecond    goaliyun.String
	BytesHitRate      goaliyun.String
	RequestHitRate    goaliyun.String
	AverageObjectSize goaliyun.String
}

type DescribeCdnMonitorDataCDNMonitorDataList []DescribeCdnMonitorDataCDNMonitorData

func (list *DescribeCdnMonitorDataCDNMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnMonitorDataCDNMonitorData)
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
