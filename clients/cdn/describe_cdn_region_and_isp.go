package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCdnRegionAndIspRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeCdnRegionAndIspRequest) Invoke(client goaliyun.Client) (*DescribeCdnRegionAndIspResponse, error) {
	resp := &DescribeCdnRegionAndIspResponse{}
	err := client.Request("cdn", "DescribeCdnRegionAndIsp", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCdnRegionAndIspResponse struct {
	RequestId goaliyun.String
	Regions   DescribeCdnRegionAndIspRegionList
	Isps      DescribeCdnRegionAndIspIspList
}

type DescribeCdnRegionAndIspRegion struct {
	NameZh goaliyun.String
	NameEn goaliyun.String
}

type DescribeCdnRegionAndIspIsp struct {
	NameZh goaliyun.String
	NameEn goaliyun.String
}

type DescribeCdnRegionAndIspRegionList []DescribeCdnRegionAndIspRegion

func (list *DescribeCdnRegionAndIspRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnRegionAndIspRegion)
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

type DescribeCdnRegionAndIspIspList []DescribeCdnRegionAndIspIsp

func (list *DescribeCdnRegionAndIspIspList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnRegionAndIspIsp)
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
