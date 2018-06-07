package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainDownstreamBpsOfEdgeRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainDownstreamBpsOfEdgeRequest) Invoke(client goaliyun.Client) (*DescribeDomainDownstreamBpsOfEdgeResponse, error) {
	resp := &DescribeDomainDownstreamBpsOfEdgeResponse{}
	err := client.Request("cdn", "DescribeDomainDownstreamBpsOfEdge", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainDownstreamBpsOfEdgeResponse struct {
	RequestId goaliyun.String
	BpsDatas  DescribeDomainDownstreamBpsOfEdgeDomainBpsModelList
}

type DescribeDomainDownstreamBpsOfEdgeDomainBpsModel struct {
	Time goaliyun.String
	Bps  goaliyun.Float
}

type DescribeDomainDownstreamBpsOfEdgeDomainBpsModelList []DescribeDomainDownstreamBpsOfEdgeDomainBpsModel

func (list *DescribeDomainDownstreamBpsOfEdgeDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainDownstreamBpsOfEdgeDomainBpsModel)
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
