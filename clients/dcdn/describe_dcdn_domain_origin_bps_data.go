package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainOriginBpsDataRequest struct {
	StartTime  string `position:"Query" name:"StartTime"`
	FixTimeGap string `position:"Query" name:"FixTimeGap"`
	TimeMerge  string `position:"Query" name:"TimeMerge"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Interval   string `position:"Query" name:"Interval"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnDomainOriginBpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainOriginBpsDataResponse, error) {
	resp := &DescribeDcdnDomainOriginBpsDataResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainOriginBpsData", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainOriginBpsDataResponse struct {
	RequestId                goaliyun.String
	DomainName               goaliyun.String
	StartTime                goaliyun.String
	EndTime                  goaliyun.String
	DataInterval             goaliyun.String
	OriginBpsDataPerInterval DescribeDcdnDomainOriginBpsDataDataModuleList
}

type DescribeDcdnDomainOriginBpsDataDataModule struct {
	TimeStamp             goaliyun.String
	OriginBps             goaliyun.Float
	DynamicHttpOriginBps  goaliyun.Float
	DynamicHttpsOriginBps goaliyun.Float
	StaticHttpOriginBps   goaliyun.Float
	StaticHttpsOriginBps  goaliyun.Float
}

type DescribeDcdnDomainOriginBpsDataDataModuleList []DescribeDcdnDomainOriginBpsDataDataModule

func (list *DescribeDcdnDomainOriginBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainOriginBpsDataDataModule)
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
