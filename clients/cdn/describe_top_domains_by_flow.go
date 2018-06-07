package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeTopDomainsByFlowRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Limit         int64  `position:"Query" name:"Limit"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeTopDomainsByFlowRequest) Invoke(client goaliyun.Client) (*DescribeTopDomainsByFlowResponse, error) {
	resp := &DescribeTopDomainsByFlowResponse{}
	err := client.Request("cdn", "DescribeTopDomainsByFlow", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTopDomainsByFlowResponse struct {
	RequestId         goaliyun.String
	StartTime         goaliyun.String
	EndTime           goaliyun.String
	DomainCount       goaliyun.Integer
	DomainOnlineCount goaliyun.Integer
	TopDomains        DescribeTopDomainsByFlowTopDomainList
}

type DescribeTopDomainsByFlowTopDomain struct {
	DomainName     goaliyun.String
	Rank           goaliyun.Integer
	TotalTraffic   goaliyun.String
	TrafficPercent goaliyun.String
	MaxBps         goaliyun.Integer
	MaxBpsTime     goaliyun.String
	TotalAccess    goaliyun.Integer
}

type DescribeTopDomainsByFlowTopDomainList []DescribeTopDomainsByFlowTopDomain

func (list *DescribeTopDomainsByFlowTopDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTopDomainsByFlowTopDomain)
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
