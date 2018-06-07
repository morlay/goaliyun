package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type InsertSubScoreForApiRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *InsertSubScoreForApiRequest) Invoke(client goaliyun.Client) (*InsertSubScoreForApiResponse, error) {
	resp := &InsertSubScoreForApiResponse{}
	err := client.Request("qualitycheck", "InsertSubScoreForApi", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InsertSubScoreForApiResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      InsertSubScoreForApiData
}

type InsertSubScoreForApiData struct {
	ScoreSubId   goaliyun.Integer
	ScoreSubName goaliyun.String
}
