package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainAverageResponseTimeRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TimeMerge     string `position:"Query" name:"TimeMerge"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainAverageResponseTimeRequest) Invoke(client goaliyun.Client) (*DescribeDomainAverageResponseTimeResponse, error) {
	resp := &DescribeDomainAverageResponseTimeResponse{}
	err := client.Request("cdn", "DescribeDomainAverageResponseTime", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainAverageResponseTimeResponse struct {
	RequestId        goaliyun.String
	DomainName       goaliyun.String
	DataInterval     goaliyun.String
	StartTime        goaliyun.String
	EndTime          goaliyun.String
	AvgRTPerInterval DescribeDomainAverageResponseTimeDataModuleList
}

type DescribeDomainAverageResponseTimeDataModule struct {
	TimeStamp goaliyun.String
	Value     goaliyun.String
}

type DescribeDomainAverageResponseTimeDataModuleList []DescribeDomainAverageResponseTimeDataModule

func (list *DescribeDomainAverageResponseTimeDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainAverageResponseTimeDataModule)
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
