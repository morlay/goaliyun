package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveSpecificDomainMappingRequest struct {
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveSpecificDomainMappingRequest) Invoke(client goaliyun.Client) (*DescribeLiveSpecificDomainMappingResponse, error) {
	resp := &DescribeLiveSpecificDomainMappingResponse{}
	err := client.Request("cdn", "DescribeLiveSpecificDomainMapping", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveSpecificDomainMappingResponse struct {
	RequestId           goaliyun.String
	DomainMappingModels DescribeLiveSpecificDomainMappingDomainMappingModelList
}

type DescribeLiveSpecificDomainMappingDomainMappingModel struct {
	PushDomain goaliyun.String
	PullDomain goaliyun.String
}

type DescribeLiveSpecificDomainMappingDomainMappingModelList []DescribeLiveSpecificDomainMappingDomainMappingModel

func (list *DescribeLiveSpecificDomainMappingDomainMappingModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveSpecificDomainMappingDomainMappingModel)
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
