package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainReqHitRateDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainReqHitRateDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainReqHitRateDataResponse, error) {
	resp := &DescribeDomainReqHitRateDataResponse{}
	err := client.Request("cdn", "DescribeDomainReqHitRateData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainReqHitRateDataResponse struct {
	RequestId          goaliyun.String
	DomainName         goaliyun.String
	DataInterval       goaliyun.String
	StartTime          goaliyun.String
	EndTime            goaliyun.String
	ReqHitRateInterval DescribeDomainReqHitRateDataDataModuleList
}

type DescribeDomainReqHitRateDataDataModule struct {
	TimeStamp goaliyun.String
	Value     goaliyun.String
}

type DescribeDomainReqHitRateDataDataModuleList []DescribeDomainReqHitRateDataDataModule

func (list *DescribeDomainReqHitRateDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainReqHitRateDataDataModule)
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
