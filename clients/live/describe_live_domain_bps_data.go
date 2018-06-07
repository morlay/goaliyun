package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveDomainBpsDataRequest struct {
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	Interval       string `position:"Query" name:"Interval"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveDomainBpsDataRequest) Invoke(client goaliyun.Client) (*DescribeLiveDomainBpsDataResponse, error) {
	resp := &DescribeLiveDomainBpsDataResponse{}
	err := client.Request("live", "DescribeLiveDomainBpsData", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveDomainBpsDataResponse struct {
	RequestId          goaliyun.String
	DomainName         goaliyun.String
	StartTime          goaliyun.String
	EndTime            goaliyun.String
	DataInterval       goaliyun.String
	BpsDataPerInterval DescribeLiveDomainBpsDataDataModuleList
}

type DescribeLiveDomainBpsDataDataModule struct {
	TimeStamp     goaliyun.String
	BpsValue      goaliyun.String
	HttpBpsValue  goaliyun.String
	HttpsBpsValue goaliyun.String
}

type DescribeLiveDomainBpsDataDataModuleList []DescribeLiveDomainBpsDataDataModule

func (list *DescribeLiveDomainBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveDomainBpsDataDataModule)
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
