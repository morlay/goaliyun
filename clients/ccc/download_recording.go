package ccc

import (
	"github.com/morlay/goaliyun"
)

type DownloadRecordingRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	FileName   string `position:"Query" name:"FileName"`
	Channel    string `position:"Query" name:"Channel"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DownloadRecordingRequest) Invoke(client goaliyun.Client) (*DownloadRecordingResponse, error) {
	resp := &DownloadRecordingResponse{}
	err := client.Request("ccc", "DownloadRecording", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DownloadRecordingResponse struct {
	RequestId          goaliyun.String
	Success            bool
	Code               goaliyun.String
	Message            goaliyun.String
	HttpStatusCode     goaliyun.Integer
	MediaDownloadParam DownloadRecordingMediaDownloadParam
}

type DownloadRecordingMediaDownloadParam struct {
	SignatureUrl goaliyun.String
	FileName     goaliyun.String
}
