package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type DeleteScoreForApiRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteScoreForApiRequest) Invoke(client goaliyun.Client) (*DeleteScoreForApiResponse, error) {
	resp := &DeleteScoreForApiResponse{}
	err := client.Request("qualitycheck", "DeleteScoreForApi", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteScoreForApiResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
