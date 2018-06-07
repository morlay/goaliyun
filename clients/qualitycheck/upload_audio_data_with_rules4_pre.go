package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UploadAudioDataWithRules4PreRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UploadAudioDataWithRules4PreRequest) Invoke(client goaliyun.Client) (*UploadAudioDataWithRules4PreResponse, error) {
	resp := &UploadAudioDataWithRules4PreResponse{}
	err := client.Request("qualitycheck", "UploadAudioDataWithRules4Pre", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadAudioDataWithRules4PreResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
