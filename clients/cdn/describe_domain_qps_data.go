package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainQpsDataRequest struct {
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

func (req *DescribeDomainQpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainQpsDataResponse, error) {
	resp := &DescribeDomainQpsDataResponse{}
	err := client.Request("cdn", "DescribeDomainQpsData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainQpsDataResponse struct {
	RequestId       goaliyun.String
	DomainName      goaliyun.String
	DataInterval    goaliyun.String
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	QpsDataInterval DescribeDomainQpsDataDataModuleList
}

type DescribeDomainQpsDataDataModule struct {
	TimeStamp            goaliyun.String
	Value                goaliyun.String
	DomesticValue        goaliyun.String
	OverseasValue        goaliyun.String
	AccValue             goaliyun.String
	AccDomesticValue     goaliyun.String
	AccOverseasValue     goaliyun.String
	DynamicValue         goaliyun.String
	DynamicDomesticValue goaliyun.String
	DynamicOverseasValue goaliyun.String
	StaticValue          goaliyun.String
	StaticDomesticValue  goaliyun.String
	StaticOverseasValue  goaliyun.String
}

type DescribeDomainQpsDataDataModuleList []DescribeDomainQpsDataDataModule

func (list *DescribeDomainQpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainQpsDataDataModule)
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
