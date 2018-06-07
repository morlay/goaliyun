package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UpdateSubScoreForApiRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateSubScoreForApiRequest) Invoke(client goaliyun.Client) (*UpdateSubScoreForApiResponse, error) {
	resp := &UpdateSubScoreForApiResponse{}
	err := client.Request("qualitycheck", "UpdateSubScoreForApi", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateSubScoreForApiResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
