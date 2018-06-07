package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveDomainTrafficDataRequest struct {
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	Interval       string `position:"Query" name:"Interval"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveDomainTrafficDataRequest) Invoke(client goaliyun.Client) (*DescribeLiveDomainTrafficDataResponse, error) {
	resp := &DescribeLiveDomainTrafficDataResponse{}
	err := client.Request("live", "DescribeLiveDomainTrafficData", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveDomainTrafficDataResponse struct {
	RequestId              goaliyun.String
	DomainName             goaliyun.String
	StartTime              goaliyun.String
	EndTime                goaliyun.String
	DataInterval           goaliyun.String
	TrafficDataPerInterval DescribeLiveDomainTrafficDataDataModuleList
}

type DescribeLiveDomainTrafficDataDataModule struct {
	TimeStamp         goaliyun.String
	TrafficValue      goaliyun.String
	HttpTrafficValue  goaliyun.String
	HttpsTrafficValue goaliyun.String
}

type DescribeLiveDomainTrafficDataDataModuleList []DescribeLiveDomainTrafficDataDataModule

func (list *DescribeLiveDomainTrafficDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveDomainTrafficDataDataModule)
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
