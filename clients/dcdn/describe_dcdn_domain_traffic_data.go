package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainTrafficDataRequest struct {
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	FixTimeGap     string `position:"Query" name:"FixTimeGap"`
	TimeMerge      string `position:"Query" name:"TimeMerge"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	Interval       string `position:"Query" name:"Interval"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnDomainTrafficDataRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainTrafficDataResponse, error) {
	resp := &DescribeDcdnDomainTrafficDataResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainTrafficData", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainTrafficDataResponse struct {
	RequestId              goaliyun.String
	DomainName             goaliyun.String
	StartTime              goaliyun.String
	EndTime                goaliyun.String
	DataInterval           goaliyun.String
	TrafficDataPerInterval DescribeDcdnDomainTrafficDataDataModuleList
}

type DescribeDcdnDomainTrafficDataDataModule struct {
	TimeStamp           goaliyun.String
	Traffic             goaliyun.Float
	DynamicHttpTraffic  goaliyun.Float
	DynamicHttpsTraffic goaliyun.Float
	StaticHttpTraffic   goaliyun.Float
	StaticHttpsTraffic  goaliyun.Float
}

type DescribeDcdnDomainTrafficDataDataModuleList []DescribeDcdnDomainTrafficDataDataModule

func (list *DescribeDcdnDomainTrafficDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainTrafficDataDataModule)
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
