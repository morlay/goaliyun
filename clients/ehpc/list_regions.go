package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListRegionsRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListRegionsRequest) Invoke(client goaliyun.Client) (*ListRegionsResponse, error) {
	resp := &ListRegionsResponse{}
	err := client.Request("ehpc", "ListRegions", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListRegionsResponse struct {
	RequestId goaliyun.String
	Regions   ListRegionsRegionInfoList
}

type ListRegionsRegionInfo struct {
	RegionId  goaliyun.String
	LocalName goaliyun.String
}

type ListRegionsRegionInfoList []ListRegionsRegionInfo

func (list *ListRegionsRegionInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRegionsRegionInfo)
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
