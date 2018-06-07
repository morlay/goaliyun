package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type DeleteAsrVocabRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteAsrVocabRequest) Invoke(client goaliyun.Client) (*DeleteAsrVocabResponse, error) {
	resp := &DeleteAsrVocabResponse{}
	err := client.Request("qualitycheck", "DeleteAsrVocab", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAsrVocabResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
