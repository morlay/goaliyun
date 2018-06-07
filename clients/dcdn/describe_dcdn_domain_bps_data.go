package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainBpsDataRequest struct {
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

func (req *DescribeDcdnDomainBpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainBpsDataResponse, error) {
	resp := &DescribeDcdnDomainBpsDataResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainBpsData", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainBpsDataResponse struct {
	RequestId          goaliyun.String
	DomainName         goaliyun.String
	StartTime          goaliyun.String
	EndTime            goaliyun.String
	DataInterval       goaliyun.String
	BpsDataPerInterval DescribeDcdnDomainBpsDataDataModuleList
}

type DescribeDcdnDomainBpsDataDataModule struct {
	TimeStamp       goaliyun.String
	Bps             goaliyun.Float
	DynamicHttpBps  goaliyun.Float
	DynamicHttpsBps goaliyun.Float
	StaticHttpBps   goaliyun.Float
	StaticHttpsBps  goaliyun.Float
}

type DescribeDcdnDomainBpsDataDataModuleList []DescribeDcdnDomainBpsDataDataModule

func (list *DescribeDcdnDomainBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainBpsDataDataModule)
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
