package cbn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCenRegionDomainRouteEntriesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CenRegionId          string `position:"Query" name:"CenRegionId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCenRegionDomainRouteEntriesRequest) Invoke(client goaliyun.Client) (*DescribeCenRegionDomainRouteEntriesResponse, error) {
	resp := &DescribeCenRegionDomainRouteEntriesResponse{}
	err := client.Request("cbn", "DescribeCenRegionDomainRouteEntries", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCenRegionDomainRouteEntriesResponse struct {
	RequestId       goaliyun.String
	PageNumber      goaliyun.Integer
	TotalCount      goaliyun.Integer
	PageSize        goaliyun.Integer
	CenRouteEntries DescribeCenRegionDomainRouteEntriesCenRouteEntryList
}

type DescribeCenRegionDomainRouteEntriesCenRouteEntry struct {
	DestinationCidrBlock goaliyun.String
	Type                 goaliyun.String
	NextHopInstanceId    goaliyun.String
	NextHopType          goaliyun.String
	NextHopRegionId      goaliyun.String
}

type DescribeCenRegionDomainRouteEntriesCenRouteEntryList []DescribeCenRegionDomainRouteEntriesCenRouteEntry

func (list *DescribeCenRegionDomainRouteEntriesCenRouteEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenRegionDomainRouteEntriesCenRouteEntry)
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
