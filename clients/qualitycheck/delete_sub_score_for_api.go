package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type DeleteSubScoreForApiRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteSubScoreForApiRequest) Invoke(client goaliyun.Client) (*DeleteSubScoreForApiResponse, error) {
	resp := &DeleteSubScoreForApiResponse{}
	err := client.Request("qualitycheck", "DeleteSubScoreForApi", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteSubScoreForApiResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
