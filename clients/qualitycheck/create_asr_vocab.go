package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type CreateAsrVocabRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CreateAsrVocabRequest) Invoke(client goaliyun.Client) (*CreateAsrVocabResponse, error) {
	resp := &CreateAsrVocabResponse{}
	err := client.Request("qualitycheck", "CreateAsrVocab", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAsrVocabResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
