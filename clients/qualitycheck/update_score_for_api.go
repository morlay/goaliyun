package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UpdateScoreForApiRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateScoreForApiRequest) Invoke(client goaliyun.Client) (*UpdateScoreForApiResponse, error) {
	resp := &UpdateScoreForApiResponse{}
	err := client.Request("qualitycheck", "UpdateScoreForApi", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateScoreForApiResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
