package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDnsProductInstancesRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	VersionCode  string `position:"Query" name:"VersionCode"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDnsProductInstancesRequest) Invoke(client goaliyun.Client) (*DescribeDnsProductInstancesResponse, error) {
	resp := &DescribeDnsProductInstancesResponse{}
	err := client.Request("alidns", "DescribeDnsProductInstances", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDnsProductInstancesResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	DnsProducts DescribeDnsProductInstancesDnsProductList
}

type DescribeDnsProductInstancesDnsProduct struct {
	InstanceId            goaliyun.String
	VersionCode           goaliyun.String
	VersionName           goaliyun.String
	StartTime             goaliyun.String
	EndTime               goaliyun.String
	StartTimestamp        goaliyun.Integer
	EndTimestamp          goaliyun.Integer
	Domain                goaliyun.String
	BindCount             goaliyun.Integer
	BindUsedCount         goaliyun.Integer
	TTLMinValue           goaliyun.Integer
	SubDomainLevel        goaliyun.Integer
	DnsSLBCount           goaliyun.Integer
	URLForwardCount       goaliyun.Integer
	DDosDefendFlow        goaliyun.Integer
	DDosDefendQuery       goaliyun.Integer
	OverseaDDosDefendFlow goaliyun.Integer
	SearchEngineLines     goaliyun.String
	ISPLines              goaliyun.String
	ISPRegionLines        goaliyun.String
	OverseaLine           goaliyun.String
	MonitorNodeCount      goaliyun.Integer
	MonitorFrequency      goaliyun.Integer
	MonitorTaskCount      goaliyun.Integer
}

type DescribeDnsProductInstancesDnsProductList []DescribeDnsProductInstancesDnsProduct

func (list *DescribeDnsProductInstancesDnsProductList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDnsProductInstancesDnsProduct)
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
