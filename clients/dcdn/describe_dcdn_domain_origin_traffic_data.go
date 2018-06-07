package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainOriginTrafficDataRequest struct {
	StartTime  string `position:"Query" name:"StartTime"`
	FixTimeGap string `position:"Query" name:"FixTimeGap"`
	TimeMerge  string `position:"Query" name:"TimeMerge"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Interval   string `position:"Query" name:"Interval"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnDomainOriginTrafficDataRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainOriginTrafficDataResponse, error) {
	resp := &DescribeDcdnDomainOriginTrafficDataResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainOriginTrafficData", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainOriginTrafficDataResponse struct {
	RequestId                    goaliyun.String
	DomainName                   goaliyun.String
	StartTime                    goaliyun.String
	EndTime                      goaliyun.String
	DataInterval                 goaliyun.String
	OriginTrafficDataPerInterval DescribeDcdnDomainOriginTrafficDataDataModuleList
}

type DescribeDcdnDomainOriginTrafficDataDataModule struct {
	TimeStamp                 goaliyun.String
	OriginTraffic             goaliyun.Float
	DynamicHttpOriginTraffic  goaliyun.Float
	DynamicHttpsOriginTraffic goaliyun.Float
	StaticHttpOriginTraffic   goaliyun.Float
	StaticHttpsOriginTraffic  goaliyun.Float
}

type DescribeDcdnDomainOriginTrafficDataDataModuleList []DescribeDcdnDomainOriginTrafficDataDataModule

func (list *DescribeDcdnDomainOriginTrafficDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainOriginTrafficDataDataModule)
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
