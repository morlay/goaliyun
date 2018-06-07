package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SubmitMediaInfoJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitMediaInfoJobRequest) Invoke(client goaliyun.Client) (*SubmitMediaInfoJobResponse, error) {
	resp := &SubmitMediaInfoJobResponse{}
	err := client.Request("mts", "SubmitMediaInfoJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitMediaInfoJobResponse struct {
	RequestId    goaliyun.String
	MediaInfoJob SubmitMediaInfoJobMediaInfoJob
}

type SubmitMediaInfoJobMediaInfoJob struct {
	JobId            goaliyun.String
	UserData         goaliyun.String
	PipelineId       goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	CreationTime     goaliyun.String
	Input            SubmitMediaInfoJobInput
	Properties       SubmitMediaInfoJobProperties
	MNSMessageResult SubmitMediaInfoJobMNSMessageResult
}

type SubmitMediaInfoJobInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type SubmitMediaInfoJobProperties struct {
	Width      goaliyun.String
	Height     goaliyun.String
	Bitrate    goaliyun.String
	Duration   goaliyun.String
	Fps        goaliyun.String
	FileSize   goaliyun.String
	FileFormat goaliyun.String
	Streams    SubmitMediaInfoJobStreams
	Format     SubmitMediaInfoJobFormat
}

type SubmitMediaInfoJobStreams struct {
	VideoStreamList    SubmitMediaInfoJobVideoStreamList
	AudioStreamList    SubmitMediaInfoJobAudioStreamList
	SubtitleStreamList SubmitMediaInfoJobSubtitleStreamList
}

type SubmitMediaInfoJobVideoStream struct {
	Index          goaliyun.String
	CodecName      goaliyun.String
	CodecLongName  goaliyun.String
	Profile        goaliyun.String
	CodecTimeBase  goaliyun.String
	CodecTagString goaliyun.String
	CodecTag       goaliyun.String
	Width          goaliyun.String
	Height         goaliyun.String
	HasBFrames     goaliyun.String
	Sar            goaliyun.String
	Dar            goaliyun.String
	PixFmt         goaliyun.String
	Level          goaliyun.String
	Fps            goaliyun.String
	AvgFPS         goaliyun.String
	Timebase       goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Bitrate        goaliyun.String
	NumFrames      goaliyun.String
	Lang           goaliyun.String
	Rotate         goaliyun.String
	NetworkCost    SubmitMediaInfoJobNetworkCost
}

type SubmitMediaInfoJobNetworkCost struct {
	PreloadTime   goaliyun.String
	CostBandwidth goaliyun.String
	AvgBitrate    goaliyun.String
}

type SubmitMediaInfoJobAudioStream struct {
	Index          goaliyun.String
	CodecName      goaliyun.String
	CodecTimeBase  goaliyun.String
	CodecLongName  goaliyun.String
	CodecTagString goaliyun.String
	CodecTag       goaliyun.String
	SampleFmt      goaliyun.String
	Samplerate     goaliyun.String
	Channels       goaliyun.String
	ChannelLayout  goaliyun.String
	Timebase       goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Bitrate        goaliyun.String
	NumFrames      goaliyun.String
	Lang           goaliyun.String
}

type SubmitMediaInfoJobSubtitleStream struct {
	Index goaliyun.String
	Lang  goaliyun.String
}

type SubmitMediaInfoJobFormat struct {
	NumStreams     goaliyun.String
	NumPrograms    goaliyun.String
	FormatName     goaliyun.String
	FormatLongName goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Size           goaliyun.String
	Bitrate        goaliyun.String
}

type SubmitMediaInfoJobMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type SubmitMediaInfoJobVideoStreamList []SubmitMediaInfoJobVideoStream

func (list *SubmitMediaInfoJobVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMediaInfoJobVideoStream)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type SubmitMediaInfoJobAudioStreamList []SubmitMediaInfoJobAudioStream

func (list *SubmitMediaInfoJobAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMediaInfoJobAudioStream)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type SubmitMediaInfoJobSubtitleStreamList []SubmitMediaInfoJobSubtitleStream

func (list *SubmitMediaInfoJobSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMediaInfoJobSubtitleStream)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
