package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainTopUrlVisitRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	SortBy        string `position:"Query" name:"SortBy"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainTopUrlVisitRequest) Invoke(client goaliyun.Client) (*DescribeDomainTopUrlVisitResponse, error) {
	resp := &DescribeDomainTopUrlVisitResponse{}
	err := client.Request("cdn", "DescribeDomainTopUrlVisit", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainTopUrlVisitResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.String
	StartTime  goaliyun.String
	AllUrlList DescribeDomainTopUrlVisitUrlListList
	Url200List DescribeDomainTopUrlVisitUrlListList
	Url300List DescribeDomainTopUrlVisitUrlListList
	Url400List DescribeDomainTopUrlVisitUrlListList
	Url500List DescribeDomainTopUrlVisitUrlListList
}

type DescribeDomainTopUrlVisitUrlList struct {
	UrlDetail       goaliyun.String
	VisitData       goaliyun.String
	VisitProportion goaliyun.Float
	Flow            goaliyun.String
	FlowProportion  goaliyun.Float
}

type DescribeDomainTopUrlVisitUrlListList []DescribeDomainTopUrlVisitUrlList

func (list *DescribeDomainTopUrlVisitUrlListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainTopUrlVisitUrlList)
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
