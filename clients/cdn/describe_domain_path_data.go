package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainPathDataRequest struct {
	StartTime     string `position:"Query" name:"StartTime"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
	Path          string `position:"Query" name:"Path"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainPathDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainPathDataResponse, error) {
	resp := &DescribeDomainPathDataResponse{}
	err := client.Request("cdn", "DescribeDomainPathData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainPathDataResponse struct {
	DomainName          goaliyun.String
	StartTime           goaliyun.String
	EndTime             goaliyun.String
	PageSize            goaliyun.Integer
	PageNumber          goaliyun.Integer
	DataInterval        goaliyun.String
	TotalCount          goaliyun.Integer
	PathDataPerInterval DescribeDomainPathDataUsageDataList
}

type DescribeDomainPathDataUsageData struct {
	Traffic goaliyun.Integer
	Acc     goaliyun.Integer
	Path    goaliyun.String
	Time    goaliyun.String
}

type DescribeDomainPathDataUsageDataList []DescribeDomainPathDataUsageData

func (list *DescribeDomainPathDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainPathDataUsageData)
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
