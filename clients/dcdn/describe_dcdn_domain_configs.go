package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainConfigsRequest struct {
	FunctionNames string `position:"Query" name:"FunctionNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnDomainConfigsRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainConfigsResponse, error) {
	resp := &DescribeDcdnDomainConfigsResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainConfigs", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainConfigsResponse struct {
	RequestId     goaliyun.String
	DomainConfigs DescribeDcdnDomainConfigsDomainConfigList
}

type DescribeDcdnDomainConfigsDomainConfig struct {
	FunctionName goaliyun.String
	ConfigId     goaliyun.String
	Status       goaliyun.String
	FunctionArgs DescribeDcdnDomainConfigsFunctionArgList
}

type DescribeDcdnDomainConfigsFunctionArg struct {
	ArgName  goaliyun.String
	ArgValue goaliyun.String
}

type DescribeDcdnDomainConfigsDomainConfigList []DescribeDcdnDomainConfigsDomainConfig

func (list *DescribeDcdnDomainConfigsDomainConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainConfigsDomainConfig)
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

type DescribeDcdnDomainConfigsFunctionArgList []DescribeDcdnDomainConfigsFunctionArg

func (list *DescribeDcdnDomainConfigsFunctionArgList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainConfigsFunctionArg)
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
