package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type SaveReviewResultRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *SaveReviewResultRequest) Invoke(client goaliyun.Client) (*SaveReviewResultResponse, error) {
	resp := &SaveReviewResultResponse{}
	err := client.Request("qualitycheck", "SaveReviewResult", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveReviewResultResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
