package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCdnTypesRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeCdnTypesRequest) Invoke(client goaliyun.Client) (*DescribeCdnTypesResponse, error) {
	resp := &DescribeCdnTypesResponse{}
	err := client.Request("cdn", "DescribeCdnTypes", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCdnTypesResponse struct {
	RequestId goaliyun.String
	CdnTypes  DescribeCdnTypesCdnTypeList
}

type DescribeCdnTypesCdnType struct {
	Type goaliyun.String
	Desc goaliyun.String
}

type DescribeCdnTypesCdnTypeList []DescribeCdnTypesCdnType

func (list *DescribeCdnTypesCdnTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnTypesCdnType)
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
