package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDNSSLBSubDomainsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDNSSLBSubDomainsRequest) Invoke(client goaliyun.Client) (*DescribeDNSSLBSubDomainsResponse, error) {
	resp := &DescribeDNSSLBSubDomainsResponse{}
	err := client.Request("alidns", "DescribeDNSSLBSubDomains", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDNSSLBSubDomainsResponse struct {
	RequestId     goaliyun.String
	TotalCount    goaliyun.Integer
	PageNumber    goaliyun.Integer
	PageSize      goaliyun.Integer
	SlbSubDomains DescribeDNSSLBSubDomainsSlbSubDomainList
}

type DescribeDNSSLBSubDomainsSlbSubDomain struct {
	SubDomain   goaliyun.String
	RecordCount goaliyun.Integer
	Open        bool
}

type DescribeDNSSLBSubDomainsSlbSubDomainList []DescribeDNSSLBSubDomainsSlbSubDomain

func (list *DescribeDNSSLBSubDomainsSlbSubDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDNSSLBSubDomainsSlbSubDomain)
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
