package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPreferredEcsTypesRequest struct {
	SpotStrategy       string `position:"Query" name:"SpotStrategy"`
	ZoneId             string `position:"Query" name:"ZoneId"`
	InstanceChargeType string `position:"Query" name:"InstanceChargeType"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *ListPreferredEcsTypesRequest) Invoke(client goaliyun.Client) (*ListPreferredEcsTypesResponse, error) {
	resp := &ListPreferredEcsTypesResponse{}
	err := client.Request("ehpc", "ListPreferredEcsTypes", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPreferredEcsTypesResponse struct {
	RequestId           goaliyun.String
	SupportSpotInstance bool
	Series              ListPreferredEcsTypesSeriesInfoList
}

type ListPreferredEcsTypesSeriesInfo struct {
	SeriesId   goaliyun.String
	SeriesName goaliyun.String
	Roles      ListPreferredEcsTypesRoles
}

type ListPreferredEcsTypesRoles struct {
	Manager ListPreferredEcsTypesManagerList
	Compute ListPreferredEcsTypesComputeList
	Login   ListPreferredEcsTypesLoginList
}

type ListPreferredEcsTypesSeriesInfoList []ListPreferredEcsTypesSeriesInfo

func (list *ListPreferredEcsTypesSeriesInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPreferredEcsTypesSeriesInfo)
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

type ListPreferredEcsTypesManagerList []goaliyun.String

func (list *ListPreferredEcsTypesManagerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type ListPreferredEcsTypesComputeList []goaliyun.String

func (list *ListPreferredEcsTypesComputeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type ListPreferredEcsTypesLoginList []goaliyun.String

func (list *ListPreferredEcsTypesLoginList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
