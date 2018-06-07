package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryMediaInfoJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MediaInfoJobIds      string `position:"Query" name:"MediaInfoJobIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryMediaInfoJobListRequest) Invoke(client goaliyun.Client) (*QueryMediaInfoJobListResponse, error) {
	resp := &QueryMediaInfoJobListResponse{}
	err := client.Request("mts", "QueryMediaInfoJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMediaInfoJobListResponse struct {
	RequestId               goaliyun.String
	MediaInfoJobList        QueryMediaInfoJobListMediaInfoJobList
	NonExistMediaInfoJobIds QueryMediaInfoJobListNonExistMediaInfoJobIdList
}

type QueryMediaInfoJobListMediaInfoJob struct {
	JobId            goaliyun.String
	UserData         goaliyun.String
	PipelineId       goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	CreationTime     goaliyun.String
	Input            QueryMediaInfoJobListInput
	Properties       QueryMediaInfoJobListProperties
	MNSMessageResult QueryMediaInfoJobListMNSMessageResult
}

type QueryMediaInfoJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryMediaInfoJobListProperties struct {
	Width      goaliyun.String
	Height     goaliyun.String
	Bitrate    goaliyun.String
	Duration   goaliyun.String
	Fps        goaliyun.String
	FileSize   goaliyun.String
	FileFormat goaliyun.String
	Streams    QueryMediaInfoJobListStreams
	Format     QueryMediaInfoJobListFormat
}

type QueryMediaInfoJobListStreams struct {
	VideoStreamList    QueryMediaInfoJobListVideoStreamList
	AudioStreamList    QueryMediaInfoJobListAudioStreamList
	SubtitleStreamList QueryMediaInfoJobListSubtitleStreamList
}

type QueryMediaInfoJobListVideoStream struct {
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
	NetworkCost    QueryMediaInfoJobListNetworkCost
}

type QueryMediaInfoJobListNetworkCost struct {
	PreloadTime   goaliyun.String
	CostBandwidth goaliyun.String
	AvgBitrate    goaliyun.String
}

type QueryMediaInfoJobListAudioStream struct {
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

type QueryMediaInfoJobListSubtitleStream struct {
	Index goaliyun.String
	Lang  goaliyun.String
}

type QueryMediaInfoJobListFormat struct {
	NumStreams     goaliyun.String
	NumPrograms    goaliyun.String
	FormatName     goaliyun.String
	FormatLongName goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Size           goaliyun.String
	Bitrate        goaliyun.String
}

type QueryMediaInfoJobListMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type QueryMediaInfoJobListMediaInfoJobList []QueryMediaInfoJobListMediaInfoJob

func (list *QueryMediaInfoJobListMediaInfoJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListMediaInfoJob)
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

type QueryMediaInfoJobListNonExistMediaInfoJobIdList []goaliyun.String

func (list *QueryMediaInfoJobListNonExistMediaInfoJobIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type QueryMediaInfoJobListVideoStreamList []QueryMediaInfoJobListVideoStream

func (list *QueryMediaInfoJobListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListVideoStream)
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

type QueryMediaInfoJobListAudioStreamList []QueryMediaInfoJobListAudioStream

func (list *QueryMediaInfoJobListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListAudioStream)
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

type QueryMediaInfoJobListSubtitleStreamList []QueryMediaInfoJobListSubtitleStream

func (list *QueryMediaInfoJobListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListSubtitleStream)
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
