package dyvmsapi

import (
	"github.com/morlay/goaliyun"
)

type SingleCallByVoiceRequest struct {
	Volume               int64  `position:"Query" name:"Volume"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CalledNumber         string `position:"Query" name:"CalledNumber"`
	VoiceCode            string `position:"Query" name:"VoiceCode"`
	CalledShowNumber     string `position:"Query" name:"CalledShowNumber"`
	PlayTimes            int64  `position:"Query" name:"PlayTimes"`
	OutId                string `position:"Query" name:"OutId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Speed                int64  `position:"Query" name:"Speed"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SingleCallByVoiceRequest) Invoke(client goaliyun.Client) (*SingleCallByVoiceResponse, error) {
	resp := &SingleCallByVoiceResponse{}
	err := client.Request("dyvmsapi", "SingleCallByVoice", "2017-05-25", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SingleCallByVoiceResponse struct {
	RequestId goaliyun.String
	CallId    goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
}
