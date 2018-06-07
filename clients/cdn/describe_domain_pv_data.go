package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainPvDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainPvDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainPvDataResponse, error) {
	resp := &DescribeDomainPvDataResponse{}
	err := client.Request("cdn", "DescribeDomainPvData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainPvDataResponse struct {
	RequestId      goaliyun.String
	DomainName     goaliyun.String
	DataInterval   goaliyun.String
	StartTime      goaliyun.String
	EndTime        goaliyun.String
	PvDataInterval DescribeDomainPvDataUsageDataList
}

type DescribeDomainPvDataUsageData struct {
	TimeStamp goaliyun.String
	Value     goaliyun.String
}

type DescribeDomainPvDataUsageDataList []DescribeDomainPvDataUsageData

func (list *DescribeDomainPvDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainPvDataUsageData)
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
