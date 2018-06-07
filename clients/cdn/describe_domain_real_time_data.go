package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainRealTimeDataRequest struct {
	Field         string `position:"Query" name:"Field"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainRealTimeDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainRealTimeDataResponse, error) {
	resp := &DescribeDomainRealTimeDataResponse{}
	err := client.Request("cdn", "DescribeDomainRealTimeData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainRealTimeDataResponse struct {
	RequestId       goaliyun.String
	DomainName      goaliyun.String
	Field           goaliyun.String
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	DataPerInterval DescribeDomainRealTimeDataDataModuleList
}

type DescribeDomainRealTimeDataDataModule struct {
	TimeStamp goaliyun.String
	Value     goaliyun.String
}

type DescribeDomainRealTimeDataDataModuleList []DescribeDomainRealTimeDataDataModule

func (list *DescribeDomainRealTimeDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeDataDataModule)
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
