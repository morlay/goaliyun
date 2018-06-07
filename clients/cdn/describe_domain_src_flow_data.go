package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainSrcFlowDataRequest struct {
	StartTime  string `position:"Query" name:"StartTime"`
	FixTimeGap string `position:"Query" name:"FixTimeGap"`
	TimeMerge  string `position:"Query" name:"TimeMerge"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Interval   string `position:"Query" name:"Interval"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainSrcFlowDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainSrcFlowDataResponse, error) {
	resp := &DescribeDomainSrcFlowDataResponse{}
	err := client.Request("cdn", "DescribeDomainSrcFlowData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainSrcFlowDataResponse struct {
	RequestId              goaliyun.String
	DomainName             goaliyun.String
	StartTime              goaliyun.String
	EndTime                goaliyun.String
	DataInterval           goaliyun.String
	SrcFlowDataPerInterval DescribeDomainSrcFlowDataDataModuleList
}

type DescribeDomainSrcFlowDataDataModule struct {
	TimeStamp goaliyun.String
	Value     goaliyun.String
}

type DescribeDomainSrcFlowDataDataModuleList []DescribeDomainSrcFlowDataDataModule

func (list *DescribeDomainSrcFlowDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSrcFlowDataDataModule)
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
