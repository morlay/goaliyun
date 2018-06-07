package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainQpsDataRequest struct {
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	FixTimeGap     string `position:"Query" name:"FixTimeGap"`
	TimeMerge      string `position:"Query" name:"TimeMerge"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	Interval       string `position:"Query" name:"Interval"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnDomainQpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainQpsDataResponse, error) {
	resp := &DescribeDcdnDomainQpsDataResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainQpsData", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainQpsDataResponse struct {
	RequestId          goaliyun.String
	DomainName         goaliyun.String
	StartTime          goaliyun.String
	EndTime            goaliyun.String
	DataInterval       goaliyun.String
	QpsDataPerInterval DescribeDcdnDomainQpsDataDataModuleList
}

type DescribeDcdnDomainQpsDataDataModule struct {
	TimeStamp       goaliyun.String
	Qps             goaliyun.Float
	DynamicHttpQps  goaliyun.Float
	DynamicHttpsQps goaliyun.Float
	StaticHttpQps   goaliyun.Float
	StaticHttpsQps  goaliyun.Float
	Acc             goaliyun.Float
	DynamicHttpAcc  goaliyun.Float
	DynamicHttpsAcc goaliyun.Float
	StaticHttpAcc   goaliyun.Float
	StaticHttpsAcc  goaliyun.Float
}

type DescribeDcdnDomainQpsDataDataModuleList []DescribeDcdnDomainQpsDataDataModule

func (list *DescribeDcdnDomainQpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainQpsDataDataModule)
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
