package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeUserDomainsRequest struct {
	FuncFilter       string `position:"Query" name:"FuncFilter"`
	Sources          string `position:"Query" name:"Sources"`
	DomainName       string `position:"Query" name:"DomainName"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	FuncId           string `position:"Query" name:"FuncId"`
	PageNumber       int64  `position:"Query" name:"PageNumber"`
	DomainStatus     string `position:"Query" name:"DomainStatus"`
	DomainSearchType string `position:"Query" name:"DomainSearchType"`
	CheckDomainShow  string `position:"Query" name:"CheckDomainShow"`
	ResourceGroupId  string `position:"Query" name:"ResourceGroupId"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	CdnType          string `position:"Query" name:"CdnType"`
	PageSize         int64  `position:"Query" name:"PageSize"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *DescribeUserDomainsRequest) Invoke(client goaliyun.Client) (*DescribeUserDomainsResponse, error) {
	resp := &DescribeUserDomainsResponse{}
	err := client.Request("cdn", "DescribeUserDomains", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUserDomainsResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	Domains    DescribeUserDomainsPageDataList
}

type DescribeUserDomainsPageData struct {
	DomainName      goaliyun.String
	Cname           goaliyun.String
	CdnType         goaliyun.String
	DomainStatus    goaliyun.String
	GmtCreated      goaliyun.String
	GmtModified     goaliyun.String
	Description     goaliyun.String
	SourceType      goaliyun.String
	SslProtocol     goaliyun.String
	ResourceGroupId goaliyun.String
	Sandbox         goaliyun.String
	Sources         DescribeUserDomainsSourceList
}

type DescribeUserDomainsPageDataList []DescribeUserDomainsPageData

func (list *DescribeUserDomainsPageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeUserDomainsPageData)
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

type DescribeUserDomainsSourceList []goaliyun.String

func (list *DescribeUserDomainsSourceList) UnmarshalJSON(data []byte) error {
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
