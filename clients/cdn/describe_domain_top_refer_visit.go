package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainTopReferVisitRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	SortBy        string `position:"Query" name:"SortBy"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainTopReferVisitRequest) Invoke(client goaliyun.Client) (*DescribeDomainTopReferVisitResponse, error) {
	resp := &DescribeDomainTopReferVisitResponse{}
	err := client.Request("cdn", "DescribeDomainTopReferVisit", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainTopReferVisitResponse struct {
	RequestId    goaliyun.String
	DomainName   goaliyun.String
	StartTime    goaliyun.String
	TopReferList DescribeDomainTopReferVisitReferListList
}

type DescribeDomainTopReferVisitReferList struct {
	ReferDetail     goaliyun.String
	VisitData       goaliyun.String
	VisitProportion goaliyun.Float
	Flow            goaliyun.String
	FlowProportion  goaliyun.Float
}

type DescribeDomainTopReferVisitReferListList []DescribeDomainTopReferVisitReferList

func (list *DescribeDomainTopReferVisitReferListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainTopReferVisitReferList)
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
