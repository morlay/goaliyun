package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainBpsDataRequest struct {
	FixTimeGap     string `position:"Query" name:"FixTimeGap"`
	TimeMerge      string `position:"Query" name:"TimeMerge"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	DomainType     string `position:"Query" name:"DomainType"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	Interval       string `position:"Query" name:"Interval"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainBpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainBpsDataResponse, error) {
	resp := &DescribeDomainBpsDataResponse{}
	err := client.Request("cdn", "DescribeDomainBpsData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainBpsDataResponse struct {
	RequestId          goaliyun.String
	DomainName         goaliyun.String
	DataInterval       goaliyun.String
	StartTime          goaliyun.String
	EndTime            goaliyun.String
	LocationNameEn     goaliyun.String
	IspNameEn          goaliyun.String
	LocationName       goaliyun.String
	IspName            goaliyun.String
	BpsDataPerInterval DescribeDomainBpsDataDataModuleList
	SupplyBpsDatas     DescribeDomainBpsDataDataModule1List
}

type DescribeDomainBpsDataDataModule struct {
	TimeStamp            goaliyun.String
	Value                goaliyun.String
	DomesticValue        goaliyun.String
	OverseasValue        goaliyun.String
	DynamicValue         goaliyun.String
	DynamicDomesticValue goaliyun.String
	DynamicOverseasValue goaliyun.String
	StaticValue          goaliyun.String
	StaticDomesticValue  goaliyun.String
	StaticOverseasValue  goaliyun.String
	L2Value              goaliyun.String
	DomesticL2Value      goaliyun.String
	OverseasL2Value      goaliyun.String
}

type DescribeDomainBpsDataDataModule1 struct {
	TimeStamp goaliyun.String
	Value     goaliyun.String
}

type DescribeDomainBpsDataDataModuleList []DescribeDomainBpsDataDataModule

func (list *DescribeDomainBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainBpsDataDataModule)
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

type DescribeDomainBpsDataDataModule1List []DescribeDomainBpsDataDataModule1

func (list *DescribeDomainBpsDataDataModule1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainBpsDataDataModule1)
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
