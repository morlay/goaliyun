package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainFileSizeProportionDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainFileSizeProportionDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainFileSizeProportionDataResponse, error) {
	resp := &DescribeDomainFileSizeProportionDataResponse{}
	err := client.Request("cdn", "DescribeDomainFileSizeProportionData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainFileSizeProportionDataResponse struct {
	RequestId                      goaliyun.String
	DomainName                     goaliyun.String
	DataInterval                   goaliyun.String
	StartTime                      goaliyun.String
	EndTime                        goaliyun.String
	FileSizeProportionDataInterval DescribeDomainFileSizeProportionDataUsageDataList
}

type DescribeDomainFileSizeProportionDataUsageData struct {
	TimeStamp goaliyun.String
	Value     DescribeDomainFileSizeProportionDataFileSizeProportionDataList
}

type DescribeDomainFileSizeProportionDataFileSizeProportionData struct {
	FileSize   goaliyun.String
	Proportion goaliyun.String
}

type DescribeDomainFileSizeProportionDataUsageDataList []DescribeDomainFileSizeProportionDataUsageData

func (list *DescribeDomainFileSizeProportionDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainFileSizeProportionDataUsageData)
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

type DescribeDomainFileSizeProportionDataFileSizeProportionDataList []DescribeDomainFileSizeProportionDataFileSizeProportionData

func (list *DescribeDomainFileSizeProportionDataFileSizeProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainFileSizeProportionDataFileSizeProportionData)
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
