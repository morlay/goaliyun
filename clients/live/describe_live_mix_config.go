package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveMixConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveMixConfigRequest) Invoke(client goaliyun.Client) (*DescribeLiveMixConfigResponse, error) {
	resp := &DescribeLiveMixConfigResponse{}
	err := client.Request("live", "DescribeLiveMixConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveMixConfigResponse struct {
	RequestId     goaliyun.String
	MixConfigList DescribeLiveMixConfigMixConfigList
}

type DescribeLiveMixConfigMixConfig struct {
	DomainName goaliyun.String
	AppName    goaliyun.String
	Template   goaliyun.String
}

type DescribeLiveMixConfigMixConfigList []DescribeLiveMixConfigMixConfig

func (list *DescribeLiveMixConfigMixConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveMixConfigMixConfig)
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
