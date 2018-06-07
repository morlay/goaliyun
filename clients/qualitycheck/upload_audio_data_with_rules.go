package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UploadAudioDataWithRulesRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UploadAudioDataWithRulesRequest) Invoke(client goaliyun.Client) (*UploadAudioDataWithRulesResponse, error) {
	resp := &UploadAudioDataWithRulesResponse{}
	err := client.Request("qualitycheck", "UploadAudioDataWithRules", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadAudioDataWithRulesResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
