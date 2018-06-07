package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainHitRateDataRequest struct {
	StartTime  string `position:"Query" name:"StartTime"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Interval   string `position:"Query" name:"Interval"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnDomainHitRateDataRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainHitRateDataResponse, error) {
	resp := &DescribeDcdnDomainHitRateDataResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainHitRateData", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainHitRateDataResponse struct {
	RequestId          goaliyun.String
	DomainName         goaliyun.String
	StartTime          goaliyun.String
	EndTime            goaliyun.String
	DataInterval       goaliyun.String
	HitRatePerInterval DescribeDcdnDomainHitRateDataDataModuleList
}

type DescribeDcdnDomainHitRateDataDataModule struct {
	TimeStamp   goaliyun.String
	ReqHitRate  goaliyun.Float
	ByteHitRate goaliyun.Float
}

type DescribeDcdnDomainHitRateDataDataModuleList []DescribeDcdnDomainHitRateDataDataModule

func (list *DescribeDcdnDomainHitRateDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainHitRateDataDataModule)
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
