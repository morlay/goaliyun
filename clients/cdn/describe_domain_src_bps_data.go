package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainSrcBpsDataRequest struct {
	StartTime  string `position:"Query" name:"StartTime"`
	FixTimeGap string `position:"Query" name:"FixTimeGap"`
	TimeMerge  string `position:"Query" name:"TimeMerge"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Interval   string `position:"Query" name:"Interval"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainSrcBpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainSrcBpsDataResponse, error) {
	resp := &DescribeDomainSrcBpsDataResponse{}
	err := client.Request("cdn", "DescribeDomainSrcBpsData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainSrcBpsDataResponse struct {
	RequestId             goaliyun.String
	DomainName            goaliyun.String
	StartTime             goaliyun.String
	EndTime               goaliyun.String
	DataInterval          goaliyun.String
	SrcBpsDataPerInterval DescribeDomainSrcBpsDataDataModuleList
}

type DescribeDomainSrcBpsDataDataModule struct {
	TimeStamp goaliyun.String
	Value     goaliyun.String
}

type DescribeDomainSrcBpsDataDataModuleList []DescribeDomainSrcBpsDataDataModule

func (list *DescribeDomainSrcBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSrcBpsDataDataModule)
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
