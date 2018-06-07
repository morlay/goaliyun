package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainFlowDataRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	TimeMerge            string `position:"Query" name:"TimeMerge"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DomainName           string `position:"Query" name:"DomainName"`
	EndTime              string `position:"Query" name:"EndTime"`
	LocationNameEn       string `position:"Query" name:"LocationNameEn"`
	StartTime            string `position:"Query" name:"StartTime"`
	IspNameEn            string `position:"Query" name:"IspNameEn"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Interval             string `position:"Query" name:"Interval"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainFlowDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainFlowDataResponse, error) {
	resp := &DescribeDomainFlowDataResponse{}
	err := client.Request("vod", "DescribeDomainFlowData", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainFlowDataResponse struct {
	RequestId           goaliyun.String
	DomainName          goaliyun.String
	DataInterval        goaliyun.String
	StartTime           goaliyun.String
	EndTime             goaliyun.String
	FlowDataPerInterval DescribeDomainFlowDataDataModuleList
}

type DescribeDomainFlowDataDataModule struct {
	TimeStamp            goaliyun.String
	Value                goaliyun.String
	DomesticValue        goaliyun.String
	OverseasValue        goaliyun.String
	DynamicValue         goaliyun.String
	DynamicDomesticValue goaliyun.String
	DynamicOverseasValue goaliyun.String
	StaticValue          goaliyun.String
	StaticDomesticValue  goaliyun.String
	StaticOverseasValue  goaliyun.String
}

type DescribeDomainFlowDataDataModuleList []DescribeDomainFlowDataDataModule

func (list *DescribeDomainFlowDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainFlowDataDataModule)
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
