package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type BatchDescribeDomainBpsDataRequest struct {
	StartTime  string `position:"Query" name:"StartTime"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *BatchDescribeDomainBpsDataRequest) Invoke(client goaliyun.Client) (*BatchDescribeDomainBpsDataResponse, error) {
	resp := &BatchDescribeDomainBpsDataResponse{}
	err := client.Request("cdn", "BatchDescribeDomainBpsData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BatchDescribeDomainBpsDataResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	BpsDatas   BatchDescribeDomainBpsDataDataModuleList
}

type BatchDescribeDomainBpsDataDataModule struct {
	Timestamp  goaliyun.String
	L1Bps      goaliyun.Float
	L1InnerBps goaliyun.Float
	L1OutBps   goaliyun.Float
	DomainName goaliyun.String
}

type BatchDescribeDomainBpsDataDataModuleList []BatchDescribeDomainBpsDataDataModule

func (list *BatchDescribeDomainBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]BatchDescribeDomainBpsDataDataModule)
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
