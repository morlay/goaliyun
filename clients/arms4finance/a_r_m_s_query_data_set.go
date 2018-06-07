package arms4finance

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ARMSQueryDataSetRequest struct {
	Measuress     *ARMSQueryDataSetMeasuresList   `position:"Query" type:"Repeated" name:"Measures"`
	IntervalInSec int64                           `position:"Query" name:"IntervalInSec"`
	DateStr       string                          `position:"Query" name:"DateStr"`
	IsDrillDown   string                          `position:"Query" name:"IsDrillDown"`
	MinTime       int64                           `position:"Query" name:"MinTime"`
	DatasetId     int64                           `position:"Query" name:"DatasetId"`
	MaxTime       int64                           `position:"Query" name:"MaxTime"`
	Dimensionss   *ARMSQueryDataSetDimensionsList `position:"Query" type:"Repeated" name:"Dimensions"`
	RegionId      string                          `position:"Query" name:"RegionId"`
}

func (req *ARMSQueryDataSetRequest) Invoke(client goaliyun.Client) (*ARMSQueryDataSetResponse, error) {
	resp := &ARMSQueryDataSetResponse{}
	err := client.Request("arms4finance", "ARMSQueryDataSet", "2017-11-30", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ARMSQueryDataSetDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type ARMSQueryDataSetResponse struct {
	Data goaliyun.String
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
