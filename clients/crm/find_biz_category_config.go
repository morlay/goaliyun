package crm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindBizCategoryConfigRequest struct {
	KpId     int64  `position:"Query" name:"KpId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *FindBizCategoryConfigRequest) Invoke(client goaliyun.Client) (*FindBizCategoryConfigResponse, error) {
	resp := &FindBizCategoryConfigResponse{}
	err := client.Request("crm", "FindBizCategoryConfig", "2015-03-24", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindBizCategoryConfigResponse struct {
	Success       bool
	ResultCode    goaliyun.String
	ResultMessage goaliyun.String
	Data          FindBizCategoryConfigBizCategoryList
}

type FindBizCategoryConfigBizCategory struct {
	Name       goaliyun.String
	Code       goaliyun.String
	IsCheck    bool
	MainBiz    bool
	Other      goaliyun.String
	SubConfigs FindBizCategoryConfigBizSubCategoryList
}

type FindBizCategoryConfigBizSubCategory struct {
	Name    goaliyun.String
	Code    goaliyun.String
	IsCheck bool
	MainBiz bool
	Other   goaliyun.String
}

type FindBizCategoryConfigBizCategoryList []FindBizCategoryConfigBizCategory

func (list *FindBizCategoryConfigBizCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindBizCategoryConfigBizCategory)
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

type FindBizCategoryConfigBizSubCategoryList []FindBizCategoryConfigBizSubCategory

func (list *FindBizCategoryConfigBizSubCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindBizCategoryConfigBizSubCategory)
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
