package crm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryEnumConfigByTypeRequest struct {
	Type     string `position:"Query" name:"Type"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryEnumConfigByTypeRequest) Invoke(client goaliyun.Client) (*QueryEnumConfigByTypeResponse, error) {
	resp := &QueryEnumConfigByTypeResponse{}
	err := client.Request("crm", "QueryEnumConfigByType", "2015-03-24", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryEnumConfigByTypeResponse struct {
	Success       bool
	ResultCode    goaliyun.String
	ResultMessage goaliyun.String
	Data          QueryEnumConfigByTypeBizCategoryList
}

type QueryEnumConfigByTypeBizCategory struct {
	EnumName  goaliyun.String
	EnumValue goaliyun.String
}

type QueryEnumConfigByTypeBizCategoryList []QueryEnumConfigByTypeBizCategory

func (list *QueryEnumConfigByTypeBizCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEnumConfigByTypeBizCategory)
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
