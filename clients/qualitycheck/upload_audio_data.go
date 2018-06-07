package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UploadAudioDataRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UploadAudioDataRequest) Invoke(client goaliyun.Client) (*UploadAudioDataResponse, error) {
	resp := &UploadAudioDataResponse{}
	err := client.Request("qualitycheck", "UploadAudioData", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadAudioDataResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
