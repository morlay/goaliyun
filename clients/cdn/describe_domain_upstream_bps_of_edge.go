package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainUpstreamBpsOfEdgeRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainUpstreamBpsOfEdgeRequest) Invoke(client goaliyun.Client) (*DescribeDomainUpstreamBpsOfEdgeResponse, error) {
	resp := &DescribeDomainUpstreamBpsOfEdgeResponse{}
	err := client.Request("cdn", "DescribeDomainUpstreamBpsOfEdge", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainUpstreamBpsOfEdgeResponse struct {
	RequestId goaliyun.String
	BpsDatas  DescribeDomainUpstreamBpsOfEdgeDomainBpsModelList
}

type DescribeDomainUpstreamBpsOfEdgeDomainBpsModel struct {
	Time goaliyun.String
	Bps  goaliyun.Float
}

type DescribeDomainUpstreamBpsOfEdgeDomainBpsModelList []DescribeDomainUpstreamBpsOfEdgeDomainBpsModel

func (list *DescribeDomainUpstreamBpsOfEdgeDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainUpstreamBpsOfEdgeDomainBpsModel)
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
