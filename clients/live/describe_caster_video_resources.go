package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCasterVideoResourcesRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeCasterVideoResourcesRequest) Invoke(client goaliyun.Client) (*DescribeCasterVideoResourcesResponse, error) {
	resp := &DescribeCasterVideoResourcesResponse{}
	err := client.Request("live", "DescribeCasterVideoResources", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCasterVideoResourcesResponse struct {
	RequestId      goaliyun.String
	Total          goaliyun.Integer
	VideoResources DescribeCasterVideoResourcesVideoResourceList
}

type DescribeCasterVideoResourcesVideoResource struct {
	MaterialId    goaliyun.String
	ResourceId    goaliyun.String
	ResourceName  goaliyun.String
	LocationId    goaliyun.String
	LiveStreamUrl goaliyun.String
	RepeatNum     goaliyun.Integer
	VodUrl        goaliyun.String
	BeginOffset   goaliyun.Integer
	EndOffset     goaliyun.Integer
}

type DescribeCasterVideoResourcesVideoResourceList []DescribeCasterVideoResourcesVideoResource

func (list *DescribeCasterVideoResourcesVideoResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterVideoResourcesVideoResource)
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
