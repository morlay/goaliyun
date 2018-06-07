package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type AddBusinessCategoryRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *AddBusinessCategoryRequest) Invoke(client goaliyun.Client) (*AddBusinessCategoryResponse, error) {
	resp := &AddBusinessCategoryResponse{}
	err := client.Request("qualitycheck", "AddBusinessCategory", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddBusinessCategoryResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
