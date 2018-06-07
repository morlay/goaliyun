package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UploadAudioData4PreRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UploadAudioData4PreRequest) Invoke(client goaliyun.Client) (*UploadAudioData4PreResponse, error) {
	resp := &UploadAudioData4PreResponse{}
	err := client.Request("qualitycheck", "UploadAudioData4Pre", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadAudioData4PreResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
