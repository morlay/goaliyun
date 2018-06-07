package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type DeleteBusinessCategoryRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteBusinessCategoryRequest) Invoke(client goaliyun.Client) (*DeleteBusinessCategoryResponse, error) {
	resp := &DeleteBusinessCategoryResponse{}
	err := client.Request("qualitycheck", "DeleteBusinessCategory", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteBusinessCategoryResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
