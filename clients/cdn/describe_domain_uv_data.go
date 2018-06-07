package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainUvDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainUvDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainUvDataResponse, error) {
	resp := &DescribeDomainUvDataResponse{}
	err := client.Request("cdn", "DescribeDomainUvData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainUvDataResponse struct {
	RequestId      goaliyun.String
	DomainName     goaliyun.String
	DataInterval   goaliyun.String
	StartTime      goaliyun.String
	EndTime        goaliyun.String
	UvDataInterval DescribeDomainUvDataUsageDataList
}

type DescribeDomainUvDataUsageData struct {
	TimeStamp goaliyun.String
	Value     goaliyun.String
}

type DescribeDomainUvDataUsageDataList []DescribeDomainUvDataUsageData

func (list *DescribeDomainUvDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainUvDataUsageData)
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
