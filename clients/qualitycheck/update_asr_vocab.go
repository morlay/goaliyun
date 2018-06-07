package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UpdateAsrVocabRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateAsrVocabRequest) Invoke(client goaliyun.Client) (*UpdateAsrVocabResponse, error) {
	resp := &UpdateAsrVocabResponse{}
	err := client.Request("qualitycheck", "UpdateAsrVocab", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateAsrVocabResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
