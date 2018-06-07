package dyvmsapi

import (
	"github.com/morlay/goaliyun"
)

type SmartCallRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ActionCodeBreak      string `position:"Query" name:"ActionCodeBreak"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RecordFlag           string `position:"Query" name:"RecordFlag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Speed                int64  `position:"Query" name:"Speed"`
	Volume               int64  `position:"Query" name:"Volume"`
	DynamicId            string `position:"Query" name:"DynamicId"`
	CalledNumber         string `position:"Query" name:"CalledNumber"`
	VoiceCode            string `position:"Query" name:"VoiceCode"`
	MuteTime             int64  `position:"Query" name:"MuteTime"`
	CalledShowNumber     string `position:"Query" name:"CalledShowNumber"`
	OutId                string `position:"Query" name:"OutId"`
	AsrModelId           string `position:"Query" name:"AsrModelId"`
	PauseTime            int64  `position:"Query" name:"PauseTime"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SmartCallRequest) Invoke(client goaliyun.Client) (*SmartCallResponse, error) {
	resp := &SmartCallResponse{}
	err := client.Request("dyvmsapi", "SmartCall", "2017-05-25", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SmartCallResponse struct {
	RequestId goaliyun.String
	CallId    goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
}
