package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeExtensiveDomainDataRequest struct {
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	ExtensiveDomain string `position:"Query" name:"ExtensiveDomain"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	EndTime         string `position:"Query" name:"EndTime"`
	StartTime       string `position:"Query" name:"StartTime"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeExtensiveDomainDataRequest) Invoke(client goaliyun.Client) (*DescribeExtensiveDomainDataResponse, error) {
	resp := &DescribeExtensiveDomainDataResponse{}
	err := client.Request("cdn", "DescribeExtensiveDomainData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeExtensiveDomainDataResponse struct {
	RequestId       goaliyun.String
	ExtensiveDomain goaliyun.String
	DataInterval    goaliyun.String
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	PageNumber      goaliyun.String
	TotalCount      goaliyun.String
	PageSize        goaliyun.String
	DataPerInterval DescribeExtensiveDomainDataUsageDataList
}

type DescribeExtensiveDomainDataUsageData struct {
	ExactDomain goaliyun.String
	TimeStamp   goaliyun.String
	Acc         goaliyun.String
	Flow        goaliyun.String
}

type DescribeExtensiveDomainDataUsageDataList []DescribeExtensiveDomainDataUsageData

func (list *DescribeExtensiveDomainDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExtensiveDomainDataUsageData)
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
