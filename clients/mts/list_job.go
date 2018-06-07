package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobRequest struct {
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	NextPageToken              string `position:"Query" name:"NextPageToken"`
	StartOfJobCreatedTimeRange string `position:"Query" name:"StartOfJobCreatedTimeRange"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	MaximumPageSize            int64  `position:"Query" name:"MaximumPageSize"`
	State                      string `position:"Query" name:"State"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	EndOfJobCreatedTimeRange   string `position:"Query" name:"EndOfJobCreatedTimeRange"`
	PipelineId                 string `position:"Query" name:"PipelineId"`
	RegionId                   string `position:"Query" name:"RegionId"`
}

func (req *ListJobRequest) Invoke(client goaliyun.Client) (*ListJobResponse, error) {
	resp := &ListJobResponse{}
	err := client.Request("mts", "ListJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobResponse struct {
	RequestId     goaliyun.String
	NextPageToken goaliyun.String
	JobList       ListJobJobList
}

type ListJobJob struct {
	JobId            goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	Percent          goaliyun.Integer
	PipelineId       goaliyun.String
	CreationTime     goaliyun.String
	FinishTime       goaliyun.String
	Input            ListJobInput
	Output           ListJobOutput
	MNSMessageResult ListJobMNSMessageResult
}

type ListJobInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type ListJobOutput struct {
	TemplateId             goaliyun.String
	UserData               goaliyun.String
	Rotate                 goaliyun.String
	VideoStreamMap         goaliyun.String
	AudioStreamMap         goaliyun.String
	DeWatermark            goaliyun.String
	Priority               goaliyun.String
	WaterMarkConfigUrl     goaliyun.String
	MergeConfigUrl         goaliyun.String
	WaterMarkList          ListJobWaterMarkList
	MergeList              ListJobMergeList
	OpeningList            ListJobMerge1List
	TailSlateList          ListJobMerge2List
	OutputFile             ListJobOutputFile
	M3U8NonStandardSupport ListJobM3U8NonStandardSupport
	Properties             ListJobProperties
	Clip                   ListJobClip
	SuperReso              ListJobSuperReso
	SubtitleConfig         ListJobSubtitleConfig
	TransConfig            ListJobTransConfig
	MuxConfig              ListJobMuxConfig
	Audio                  ListJobAudio
	Video                  ListJobVideo
	Container              ListJobContainer
	Encryption             ListJobEncryption
}

type ListJobWaterMark struct {
	WaterMarkTemplateId goaliyun.String
	Width               goaliyun.String
	Height              goaliyun.String
	Dx                  goaliyun.String
	Dy                  goaliyun.String
	ReferPos            goaliyun.String
	Type                goaliyun.String
	InputFile           ListJobInputFile
}

type ListJobInputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type ListJobMerge struct {
	MergeURL goaliyun.String
	Start    goaliyun.String
	Duration goaliyun.String
	RoleArn  goaliyun.String
}

type ListJobMerge1 struct {
	OpenUrl goaliyun.String
	Start   goaliyun.String
	Width   goaliyun.String
	Height  goaliyun.String
}

type ListJobMerge2 struct {
	TailUrl       goaliyun.String
	Start         goaliyun.String
	BlendDuration goaliyun.String
	Width         goaliyun.String
	Height        goaliyun.String
	IsMergeAudio  bool
	BgColor       goaliyun.String
}

type ListJobOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type ListJobM3U8NonStandardSupport struct {
	TS ListJobTS
}

type ListJobTS struct {
	Md5Support  bool
	SizeSupport bool
}

type ListJobProperties struct {
	Width      goaliyun.String
	Height     goaliyun.String
	Bitrate    goaliyun.String
	Duration   goaliyun.String
	Fps        goaliyun.String
	FileSize   goaliyun.String
	FileFormat goaliyun.String
	Streams    ListJobStreams
	Format     ListJobFormat
}

type ListJobStreams struct {
	VideoStreamList    ListJobVideoStreamList
	AudioStreamList    ListJobAudioStreamList
	SubtitleStreamList ListJobSubtitleStreamList
}

type ListJobVideoStream struct {
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
	NetworkCost    ListJobNetworkCost
}

type ListJobNetworkCost struct {
	PreloadTime   goaliyun.String
	CostBandwidth goaliyun.String
	AvgBitrate    goaliyun.String
}

type ListJobAudioStream struct {
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

type ListJobSubtitleStream struct {
	Index goaliyun.String
	Lang  goaliyun.String
}

type ListJobFormat struct {
	NumStreams     goaliyun.String
	NumPrograms    goaliyun.String
	FormatName     goaliyun.String
	FormatLongName goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Size           goaliyun.String
	Bitrate        goaliyun.String
}

type ListJobClip struct {
	TimeSpan ListJobTimeSpan
}

type ListJobTimeSpan struct {
	Seek     goaliyun.String
	Duration goaliyun.String
}

type ListJobSuperReso struct {
	IsHalfSample goaliyun.String
}

type ListJobSubtitleConfig struct {
	SubtitleList    ListJobSubtitleList
	ExtSubtitleList ListJobExtSubtitleList
}

type ListJobSubtitle struct {
	Map goaliyun.String
}

type ListJobExtSubtitle struct {
	FontName goaliyun.String
	CharEnc  goaliyun.String
	Input3   ListJobInput3
}

type ListJobInput3 struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type ListJobTransConfig struct {
	TransMode               goaliyun.String
	IsCheckReso             goaliyun.String
	IsCheckResoFail         goaliyun.String
	IsCheckVideoBitrate     goaliyun.String
	IsCheckAudioBitrate     goaliyun.String
	AdjDarMethod            goaliyun.String
	IsCheckVideoBitrateFail goaliyun.String
	IsCheckAudioBitrateFail goaliyun.String
}

type ListJobMuxConfig struct {
	Segment ListJobSegment
	Gif     ListJobGif
}

type ListJobSegment struct {
	Duration goaliyun.String
}

type ListJobGif struct {
	Loop            goaliyun.String
	FinalDelay      goaliyun.String
	IsCustomPalette goaliyun.String
	DitherMode      goaliyun.String
}

type ListJobAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
	Volume     ListJobVolume
}

type ListJobVolume struct {
	Level  goaliyun.String
	Method goaliyun.String
}

type ListJobVideo struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Bitrate    goaliyun.String
	Crf        goaliyun.String
	Width      goaliyun.String
	Height     goaliyun.String
	Fps        goaliyun.String
	Gop        goaliyun.String
	Preset     goaliyun.String
	ScanMode   goaliyun.String
	Bufsize    goaliyun.String
	Maxrate    goaliyun.String
	PixFmt     goaliyun.String
	Degrain    goaliyun.String
	Qscale     goaliyun.String
	Crop       goaliyun.String
	Pad        goaliyun.String
	MaxFps     goaliyun.String
	BitrateBnd ListJobBitrateBnd
}

type ListJobBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type ListJobContainer struct {
	Format goaliyun.String
}

type ListJobEncryption struct {
	Type    goaliyun.String
	Id      goaliyun.String
	Key     goaliyun.String
	KeyUri  goaliyun.String
	KeyType goaliyun.String
	SkipCnt goaliyun.String
}

type ListJobMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type ListJobJobList []ListJobJob

func (list *ListJobJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobJob)
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

type ListJobWaterMarkList []ListJobWaterMark

func (list *ListJobWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobWaterMark)
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

type ListJobMergeList []ListJobMerge

func (list *ListJobMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobMerge)
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

type ListJobMerge1List []ListJobMerge1

func (list *ListJobMerge1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobMerge1)
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

type ListJobMerge2List []ListJobMerge2

func (list *ListJobMerge2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobMerge2)
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

type ListJobVideoStreamList []ListJobVideoStream

func (list *ListJobVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobVideoStream)
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

type ListJobAudioStreamList []ListJobAudioStream

func (list *ListJobAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobAudioStream)
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

type ListJobSubtitleStreamList []ListJobSubtitleStream

func (list *ListJobSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobSubtitleStream)
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

type ListJobSubtitleList []ListJobSubtitle

func (list *ListJobSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobSubtitle)
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

type ListJobExtSubtitleList []ListJobExtSubtitle

func (list *ListJobExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExtSubtitle)
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
