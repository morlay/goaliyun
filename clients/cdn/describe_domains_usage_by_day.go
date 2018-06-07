package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainsUsageByDayRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainsUsageByDayRequest) Invoke(client goaliyun.Client) (*DescribeDomainsUsageByDayResponse, error) {
	resp := &DescribeDomainsUsageByDayResponse{}
	err := client.Request("cdn", "DescribeDomainsUsageByDay", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainsUsageByDayResponse struct {
	RequestId    goaliyun.String
	DomainName   goaliyun.String
	DataInterval goaliyun.String
	StartTime    goaliyun.String
	EndTime      goaliyun.String
	UsageByDays  DescribeDomainsUsageByDayUsageByDayList
	UsageTotal   DescribeDomainsUsageByDayUsageTotal
}

type DescribeDomainsUsageByDayUsageByDay struct {
	TimeStamp      goaliyun.String
	Qps            goaliyun.String
	BytesHitRate   goaliyun.String
	RequestHitRate goaliyun.String
	MaxBps         goaliyun.String
	MaxBpsTime     goaliyun.String
	MaxSrcBps      goaliyun.String
	MaxSrcBpsTime  goaliyun.String
	TotalAccess    goaliyun.String
	TotalTraffic   goaliyun.String
}

type DescribeDomainsUsageByDayUsageTotal struct {
	BytesHitRate   goaliyun.String
	RequestHitRate goaliyun.String
	MaxBps         goaliyun.String
	MaxBpsTime     goaliyun.String
	MaxSrcBps      goaliyun.String
	MaxSrcBpsTime  goaliyun.String
	TotalAccess    goaliyun.String
	TotalTraffic   goaliyun.String
}

type DescribeDomainsUsageByDayUsageByDayList []DescribeDomainsUsageByDayUsageByDay

func (list *DescribeDomainsUsageByDayUsageByDayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsUsageByDayUsageByDay)
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
