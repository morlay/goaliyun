package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainHttpCodeDataRequest struct {
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

func (req *DescribeDomainHttpCodeDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainHttpCodeDataResponse, error) {
	resp := &DescribeDomainHttpCodeDataResponse{}
	err := client.Request("cdn", "DescribeDomainHttpCodeData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainHttpCodeDataResponse struct {
	RequestId    goaliyun.String
	DomainName   goaliyun.String
	DataInterval goaliyun.String
	StartTime    goaliyun.String
	EndTime      goaliyun.String
	HttpCodeData DescribeDomainHttpCodeDataUsageDataList
}

type DescribeDomainHttpCodeDataUsageData struct {
	TimeStamp goaliyun.String
	Value     DescribeDomainHttpCodeDataCodeProportionDataList
}

type DescribeDomainHttpCodeDataCodeProportionData struct {
	Code       goaliyun.String
	Proportion goaliyun.String
	Count      goaliyun.String
}

type DescribeDomainHttpCodeDataUsageDataList []DescribeDomainHttpCodeDataUsageData

func (list *DescribeDomainHttpCodeDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHttpCodeDataUsageData)
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

type DescribeDomainHttpCodeDataCodeProportionDataList []DescribeDomainHttpCodeDataCodeProportionData

func (list *DescribeDomainHttpCodeDataCodeProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHttpCodeDataCodeProportionData)
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
