package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetClusterStatusRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ItemType        string `position:"Query" name:"ItemType"`
	Interval        string `position:"Query" name:"Interval"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetClusterStatusRequest) Invoke(client goaliyun.Client) (*GetClusterStatusResponse, error) {
	resp := &GetClusterStatusResponse{}
	err := client.Request("emr", "GetClusterStatus", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetClusterStatusResponse struct {
	RequestId  goaliyun.String
	StatusList GetClusterStatusStatusList
}

type GetClusterStatusStatus struct {
	Name      goaliyun.String
	Legend    goaliyun.String
	Unit      goaliyun.String
	Series    GetClusterStatusSeriesInfoList
	LineNames GetClusterStatusLineNameList
	Times     GetClusterStatusTimeList
}

type GetClusterStatusSeriesInfo struct {
	SeriesItems GetClusterStatusSeriesItemList
}

type GetClusterStatusSeriesItem struct {
	SeriesValue goaliyun.Float
}

type GetClusterStatusStatusList []GetClusterStatusStatus

func (list *GetClusterStatusStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterStatusStatus)
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

type GetClusterStatusSeriesInfoList []GetClusterStatusSeriesInfo

func (list *GetClusterStatusSeriesInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterStatusSeriesInfo)
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

type GetClusterStatusLineNameList []goaliyun.String

func (list *GetClusterStatusLineNameList) UnmarshalJSON(data []byte) error {
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

type GetClusterStatusTimeList []goaliyun.String

func (list *GetClusterStatusTimeList) UnmarshalJSON(data []byte) error {
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

type GetClusterStatusSeriesItemList []GetClusterStatusSeriesItem

func (list *GetClusterStatusSeriesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterStatusSeriesItem)
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
