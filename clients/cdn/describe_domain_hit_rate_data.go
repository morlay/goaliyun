package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainHitRateDataRequest struct {
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	TimeMerge      string `position:"Query" name:"TimeMerge"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	Interval       string `position:"Query" name:"Interval"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainHitRateDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainHitRateDataResponse, error) {
	resp := &DescribeDomainHitRateDataResponse{}
	err := client.Request("cdn", "DescribeDomainHitRateData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainHitRateDataResponse struct {
	RequestId       goaliyun.String
	DomainName      goaliyun.String
	DataInterval    goaliyun.String
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	HitRateInterval DescribeDomainHitRateDataDataModuleList
}

type DescribeDomainHitRateDataDataModule struct {
	TimeStamp goaliyun.String
	Value     goaliyun.String
}

type DescribeDomainHitRateDataDataModuleList []DescribeDomainHitRateDataDataModule

func (list *DescribeDomainHitRateDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHitRateDataDataModule)
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
