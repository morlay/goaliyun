package arms4finance

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type WhereInDimQueryRequest struct {
	WhereInKey     string                            `position:"Query" name:"WhereInKey"`
	Measuress      *WhereInDimQueryMeasuresList      `position:"Query" type:"Repeated" name:"Measures"`
	IntervalInSec  int64                             `position:"Query" name:"IntervalInSec"`
	DateStr        string                            `position:"Query" name:"DateStr"`
	IsDrillDown    string                            `position:"Query" name:"IsDrillDown"`
	MinTime        int64                             `position:"Query" name:"MinTime"`
	DatasetId      int64                             `position:"Query" name:"DatasetId"`
	WhereInValuess *WhereInDimQueryWhereInValuesList `position:"Query" type:"Repeated" name:"WhereInValues"`
	MaxTime        int64                             `position:"Query" name:"MaxTime"`
	Dimensionss    *WhereInDimQueryDimensionsList    `position:"Query" type:"Repeated" name:"Dimensions"`
	RegionId       string                            `position:"Query" name:"RegionId"`
}

func (req *WhereInDimQueryRequest) Invoke(client goaliyun.Client) (*WhereInDimQueryResponse, error) {
	resp := &WhereInDimQueryResponse{}
	err := client.Request("arms4finance", "WhereInDimQuery", "2017-11-30", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type WhereInDimQueryDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type WhereInDimQueryResponse struct {
	Data goaliyun.String
}

type WhereInDimQueryMeasuresList []string

func (list *WhereInDimQueryMeasuresList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type WhereInDimQueryWhereInValuesList []string

func (list *WhereInDimQueryWhereInValuesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type WhereInDimQueryDimensionsList []WhereInDimQueryDimensions

func (list *WhereInDimQueryDimensionsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WhereInDimQueryDimensions)
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
