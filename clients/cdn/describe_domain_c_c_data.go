package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainCCDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainCCDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainCCDataResponse, error) {
	resp := &DescribeDomainCCDataResponse{}
	err := client.Request("cdn", "DescribeDomainCCData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainCCDataResponse struct {
	RequestId    goaliyun.String
	DomainName   goaliyun.String
	DataInterval goaliyun.String
	StartTime    goaliyun.String
	EndTime      goaliyun.String
	CCDataList   DescribeDomainCCDataCCDatasList
}

type DescribeDomainCCDataCCDatas struct {
	TimeStamp goaliyun.String
	Count     goaliyun.String
}

type DescribeDomainCCDataCCDatasList []DescribeDomainCCDataCCDatas

func (list *DescribeDomainCCDataCCDatasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCCDataCCDatas)
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
