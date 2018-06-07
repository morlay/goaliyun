package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCdnDomainConfigsRequest struct {
	FunctionNames string `position:"Query" name:"FunctionNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeCdnDomainConfigsRequest) Invoke(client goaliyun.Client) (*DescribeCdnDomainConfigsResponse, error) {
	resp := &DescribeCdnDomainConfigsResponse{}
	err := client.Request("cdn", "DescribeCdnDomainConfigs", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCdnDomainConfigsResponse struct {
	RequestId     goaliyun.String
	DomainConfigs DescribeCdnDomainConfigsDomainConfigList
}

type DescribeCdnDomainConfigsDomainConfig struct {
	FunctionName goaliyun.String
	ConfigId     goaliyun.String
	Status       goaliyun.String
	FunctionArgs DescribeCdnDomainConfigsFunctionArgList
}

type DescribeCdnDomainConfigsFunctionArg struct {
	ArgName  goaliyun.String
	ArgValue goaliyun.String
}

type DescribeCdnDomainConfigsDomainConfigList []DescribeCdnDomainConfigsDomainConfig

func (list *DescribeCdnDomainConfigsDomainConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnDomainConfigsDomainConfig)
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

type DescribeCdnDomainConfigsFunctionArgList []DescribeCdnDomainConfigsFunctionArg

func (list *DescribeCdnDomainConfigsFunctionArgList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnDomainConfigsFunctionArg)
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
