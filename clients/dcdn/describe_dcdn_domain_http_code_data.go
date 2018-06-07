package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainHttpCodeDataRequest struct {
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	Interval       string `position:"Query" name:"Interval"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnDomainHttpCodeDataRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainHttpCodeDataResponse, error) {
	resp := &DescribeDcdnDomainHttpCodeDataResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainHttpCodeData", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainHttpCodeDataResponse struct {
	RequestId       goaliyun.String
	DomainName      goaliyun.String
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	DataInterval    goaliyun.String
	DataPerInterval DescribeDcdnDomainHttpCodeDataDataModuleList
}

type DescribeDcdnDomainHttpCodeDataDataModule struct {
	TimeStamp               goaliyun.String
	HttpCodeDataPerInterval DescribeDcdnDomainHttpCodeDataHttpCodeDataModuleList
}

type DescribeDcdnDomainHttpCodeDataHttpCodeDataModule struct {
	Code       goaliyun.Integer
	Proportion goaliyun.Float
	Count      goaliyun.Float
}

type DescribeDcdnDomainHttpCodeDataDataModuleList []DescribeDcdnDomainHttpCodeDataDataModule

func (list *DescribeDcdnDomainHttpCodeDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainHttpCodeDataDataModule)
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

type DescribeDcdnDomainHttpCodeDataHttpCodeDataModuleList []DescribeDcdnDomainHttpCodeDataHttpCodeDataModule

func (list *DescribeDcdnDomainHttpCodeDataHttpCodeDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainHttpCodeDataHttpCodeDataModule)
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
