package dyvmsapi

import (
	"github.com/morlay/goaliyun"
)

type SingleCallByTtsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	TtsCode              string `position:"Query" name:"TtsCode"`
	PlayTimes            int64  `position:"Query" name:"PlayTimes"`
	TtsParam             string `position:"Query" name:"TtsParam"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Speed                int64  `position:"Query" name:"Speed"`
	Volume               int64  `position:"Query" name:"Volume"`
	CalledNumber         string `position:"Query" name:"CalledNumber"`
	CalledShowNumber     string `position:"Query" name:"CalledShowNumber"`
	OutId                string `position:"Query" name:"OutId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SingleCallByTtsRequest) Invoke(client goaliyun.Client) (*SingleCallByTtsResponse, error) {
	resp := &SingleCallByTtsResponse{}
	err := client.Request("dyvmsapi", "SingleCallByTts", "2017-05-25", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SingleCallByTtsResponse struct {
	RequestId goaliyun.String
	CallId    goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
}
