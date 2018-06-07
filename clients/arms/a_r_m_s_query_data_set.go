package arms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ARMSQueryDataSetRequest struct {
	DateStr       int64                             `position:"Query" name:"DateStr"`
	MinTime       int64                             `position:"Query" name:"MinTime"`
	ReduceTail    string                            `position:"Query" name:"ReduceTail"`
	MaxTime       int64                             `position:"Query" name:"MaxTime"`
	OptionalDimss *ARMSQueryDataSetOptionalDimsList `position:"Query" type:"Repeated" name:"OptionalDims"`
	Measuress     *ARMSQueryDataSetMeasuresList     `position:"Query" type:"Repeated" name:"Measures"`
	IntervalInSec int64                             `position:"Query" name:"IntervalInSec"`
	IsDrillDown   string                            `position:"Query" name:"IsDrillDown"`
	HungryMode    string                            `position:"Query" name:"HungryMode"`
	OrderByKey    string                            `position:"Query" name:"OrderByKey"`
	Limit         int64                             `position:"Query" name:"Limit"`
	DatasetId     int64                             `position:"Query" name:"DatasetId"`
	RequiredDimss *ARMSQueryDataSetRequiredDimsList `position:"Query" type:"Repeated" name:"RequiredDims"`
	Dimensionss   *ARMSQueryDataSetDimensionsList   `position:"Query" type:"Repeated" name:"Dimensions"`
	RegionId      string                            `position:"Query" name:"RegionId"`
}

func (req *ARMSQueryDataSetRequest) Invoke(client goaliyun.Client) (*ARMSQueryDataSetResponse, error) {
	resp := &ARMSQueryDataSetResponse{}
	err := client.Request("arms", "ARMSQueryDataSet", "2016-11-25", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ARMSQueryDataSetOptionalDims struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}

type ARMSQueryDataSetRequiredDims struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}

type ARMSQueryDataSetDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}

type ARMSQueryDataSetResponse struct {
	Data goaliyun.String
}

type ARMSQueryDataSetOptionalDimsList []ARMSQueryDataSetOptionalDims

func (list *ARMSQueryDataSetOptionalDimsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ARMSQueryDataSetOptionalDims)
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

type ARMSQueryDataSetMeasuresList []string

func (list *ARMSQueryDataSetMeasuresList) UnmarshalJSON(data []byte) error {
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

type ARMSQueryDataSetRequiredDimsList []ARMSQueryDataSetRequiredDims

func (list *ARMSQueryDataSetRequiredDimsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ARMSQueryDataSetRequiredDims)
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

type ARMSQueryDataSetDimensionsList []ARMSQueryDataSetDimensions

func (list *ARMSQueryDataSetDimensionsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ARMSQueryDataSetDimensions)
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
