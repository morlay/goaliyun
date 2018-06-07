package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetBusinessCategoryListRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetBusinessCategoryListRequest) Invoke(client goaliyun.Client) (*GetBusinessCategoryListResponse, error) {
	resp := &GetBusinessCategoryListResponse{}
	err := client.Request("qualitycheck", "GetBusinessCategoryList", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetBusinessCategoryListResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      GetBusinessCategoryListBusinessCategoryBasicInfoList
}

type GetBusinessCategoryListBusinessCategoryBasicInfo struct {
	Bid          goaliyun.Integer
	ServiceType  goaliyun.Integer
	BusinessName goaliyun.String
}

type GetBusinessCategoryListBusinessCategoryBasicInfoList []GetBusinessCategoryListBusinessCategoryBasicInfo

func (list *GetBusinessCategoryListBusinessCategoryBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetBusinessCategoryListBusinessCategoryBasicInfo)
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
