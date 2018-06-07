package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainUpstreamOfCenterRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainUpstreamOfCenterRequest) Invoke(client goaliyun.Client) (*DescribeDomainUpstreamOfCenterResponse, error) {
	resp := &DescribeDomainUpstreamOfCenterResponse{}
	err := client.Request("cdn", "DescribeDomainUpstreamOfCenter", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainUpstreamOfCenterResponse struct {
	RequestId goaliyun.String
	BpsDatas  DescribeDomainUpstreamOfCenterDomainBpsModelList
}

type DescribeDomainUpstreamOfCenterDomainBpsModel struct {
	Time goaliyun.String
	Bps  goaliyun.Float
}

type DescribeDomainUpstreamOfCenterDomainBpsModelList []DescribeDomainUpstreamOfCenterDomainBpsModel

func (list *DescribeDomainUpstreamOfCenterDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainUpstreamOfCenterDomainBpsModel)
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
