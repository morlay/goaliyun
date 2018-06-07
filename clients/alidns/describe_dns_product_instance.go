package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDnsProductInstanceRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDnsProductInstanceRequest) Invoke(client goaliyun.Client) (*DescribeDnsProductInstanceResponse, error) {
	resp := &DescribeDnsProductInstanceResponse{}
	err := client.Request("alidns", "DescribeDnsProductInstance", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDnsProductInstanceResponse struct {
	RequestId             goaliyun.String
	InstanceId            goaliyun.String
	VersionCode           goaliyun.String
	VersionName           goaliyun.String
	StartTime             goaliyun.String
	StartTimestamp        goaliyun.Integer
	EndTime               goaliyun.String
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
	DnsServers            DescribeDnsProductInstanceDnsServerList
}

type DescribeDnsProductInstanceDnsServerList []goaliyun.String

func (list *DescribeDnsProductInstanceDnsServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
