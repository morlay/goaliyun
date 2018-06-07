package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type GetAsrVocabRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetAsrVocabRequest) Invoke(client goaliyun.Client) (*GetAsrVocabResponse, error) {
	resp := &GetAsrVocabResponse{}
	err := client.Request("qualitycheck", "GetAsrVocab", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAsrVocabResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
