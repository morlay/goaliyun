package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type InsertScoreForApiRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *InsertScoreForApiRequest) Invoke(client goaliyun.Client) (*InsertScoreForApiResponse, error) {
	resp := &InsertScoreForApiResponse{}
	err := client.Request("qualitycheck", "InsertScoreForApi", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InsertScoreForApiResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      InsertScoreForApiData
}

type InsertScoreForApiData struct {
	ScoreId   goaliyun.Integer
	ScoreName goaliyun.String
}
