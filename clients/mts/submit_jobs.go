package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SubmitJobsRequest struct {
	Outputs              string `position:"Query" name:"Outputs"`
	Input                string `position:"Query" name:"Input"`
	OutputBucket         string `position:"Query" name:"OutputBucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OutputLocation       string `position:"Query" name:"OutputLocation"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitJobsRequest) Invoke(client goaliyun.Client) (*SubmitJobsResponse, error) {
	resp := &SubmitJobsResponse{}
	err := client.Request("mts", "SubmitJobs", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitJobsResponse struct {
	RequestId     goaliyun.String
	JobResultList SubmitJobsJobResultList
}

type SubmitJobsJobResult struct {
	Success bool
	Code    goaliyun.String
	Message goaliyun.String
	Job     SubmitJobsJob
}

type SubmitJobsJob struct {
	JobId            goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	Percent          goaliyun.Integer
	PipelineId       goaliyun.String
	CreationTime     goaliyun.String
	FinishTime       goaliyun.String
	Input            SubmitJobsInput
	Output           SubmitJobsOutput
	MNSMessageResult SubmitJobsMNSMessageResult
}

type SubmitJobsInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type SubmitJobsOutput struct {
	TemplateId             goaliyun.String
	UserData               goaliyun.String
	Rotate                 goaliyun.String
	VideoStreamMap         goaliyun.String
	AudioStreamMap         goaliyun.String
	DeWatermark            goaliyun.String
	Priority               goaliyun.String
	WaterMarkConfigUrl     goaliyun.String
	MergeConfigUrl         goaliyun.String
	WaterMarkList          SubmitJobsWaterMarkList
	MergeList              SubmitJobsMergeList
	OpeningList            SubmitJobsOpeningList
	TailSlateList          SubmitJobsTailSlateList
	DigiWaterMark          SubmitJobsDigiWaterMark
	OutputFile             SubmitJobsOutputFile
	M3U8NonStandardSupport SubmitJobsM3U8NonStandardSupport
	Properties             SubmitJobsProperties
	Clip                   SubmitJobsClip
	SuperReso              SubmitJobsSuperReso
	SubtitleConfig         SubmitJobsSubtitleConfig
	TransConfig            SubmitJobsTransConfig
	MuxConfig              SubmitJobsMuxConfig
	Audio                  SubmitJobsAudio
	Video                  SubmitJobsVideo
	Container              SubmitJobsContainer
	Encryption             SubmitJobsEncryption
}

type SubmitJobsWaterMark struct {
	WaterMarkTemplateId goaliyun.String
	Width               goaliyun.String
	Height              goaliyun.String
	Dx                  goaliyun.String
	Dy                  goaliyun.String
	ReferPos            goaliyun.String
	Type                goaliyun.String
	InputFile           SubmitJobsInputFile
}

type SubmitJobsInputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type SubmitJobsMerge struct {
	MergeURL goaliyun.String
	Start    goaliyun.String
	Duration goaliyun.String
	RoleArn  goaliyun.String
}

type SubmitJobsOpening struct {
	OpenUrl goaliyun.String
	Start   goaliyun.String
	Width   goaliyun.String
	Height  goaliyun.String
}

type SubmitJobsTailSlate struct {
	TailUrl       goaliyun.String
	Start         goaliyun.String
	BlendDuration goaliyun.String
	Width         goaliyun.String
	Height        goaliyun.String
	IsMergeAudio  bool
	BgColor       goaliyun.String
}

type SubmitJobsDigiWaterMark struct {
	Type       goaliyun.String
	Alpha      goaliyun.String
	InputFile1 SubmitJobsInputFile1
}

type SubmitJobsInputFile1 struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type SubmitJobsOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type SubmitJobsM3U8NonStandardSupport struct {
	TS SubmitJobsTS
}

type SubmitJobsTS struct {
	Md5Support  bool
	SizeSupport bool
}

type SubmitJobsProperties struct {
	Width      goaliyun.String
	Height     goaliyun.String
	Bitrate    goaliyun.String
	Duration   goaliyun.String
	Fps        goaliyun.String
	FileSize   goaliyun.String
	FileFormat goaliyun.String
	Streams    SubmitJobsStreams
	Format     SubmitJobsFormat
}

type SubmitJobsStreams struct {
	VideoStreamList    SubmitJobsVideoStreamList
	AudioStreamList    SubmitJobsAudioStreamList
	SubtitleStreamList SubmitJobsSubtitleStreamList
}

type SubmitJobsVideoStream struct {
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
	NetworkCost    SubmitJobsNetworkCost
}

type SubmitJobsNetworkCost struct {
	PreloadTime   goaliyun.String
	CostBandwidth goaliyun.String
	AvgBitrate    goaliyun.String
}

type SubmitJobsAudioStream struct {
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

type SubmitJobsSubtitleStream struct {
	Index goaliyun.String
	Lang  goaliyun.String
}

type SubmitJobsFormat struct {
	NumStreams     goaliyun.String
	NumPrograms    goaliyun.String
	FormatName     goaliyun.String
	FormatLongName goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Size           goaliyun.String
	Bitrate        goaliyun.String
}

type SubmitJobsClip struct {
	TimeSpan SubmitJobsTimeSpan
}

type SubmitJobsTimeSpan struct {
	Seek     goaliyun.String
	Duration goaliyun.String
}

type SubmitJobsSuperReso struct {
	IsHalfSample goaliyun.String
}

type SubmitJobsSubtitleConfig struct {
	SubtitleList    SubmitJobsSubtitleList
	ExtSubtitleList SubmitJobsExtSubtitleList
}

type SubmitJobsSubtitle struct {
	Map goaliyun.String
}

type SubmitJobsExtSubtitle struct {
	FontName goaliyun.String
	CharEnc  goaliyun.String
	Input2   SubmitJobsInput2
}

type SubmitJobsInput2 struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type SubmitJobsTransConfig struct {
	TransMode               goaliyun.String
	IsCheckReso             goaliyun.String
	IsCheckResoFail         goaliyun.String
	IsCheckVideoBitrate     goaliyun.String
	IsCheckAudioBitrate     goaliyun.String
	AdjDarMethod            goaliyun.String
	IsCheckVideoBitrateFail goaliyun.String
	IsCheckAudioBitrateFail goaliyun.String
}

type SubmitJobsMuxConfig struct {
	Segment SubmitJobsSegment
	Gif     SubmitJobsGif
}

type SubmitJobsSegment struct {
	Duration goaliyun.String
}

type SubmitJobsGif struct {
	Loop            goaliyun.String
	FinalDelay      goaliyun.String
	IsCustomPalette goaliyun.String
	DitherMode      goaliyun.String
}

type SubmitJobsAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
	Volume     SubmitJobsVolume
}

type SubmitJobsVolume struct {
	Level  goaliyun.String
	Method goaliyun.String
}

type SubmitJobsVideo struct {
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
	BitrateBnd SubmitJobsBitrateBnd
}

type SubmitJobsBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type SubmitJobsContainer struct {
	Format goaliyun.String
}

type SubmitJobsEncryption struct {
	Type    goaliyun.String
	Id      goaliyun.String
	Key     goaliyun.String
	KeyUri  goaliyun.String
	KeyType goaliyun.String
	SkipCnt goaliyun.String
}

type SubmitJobsMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type SubmitJobsJobResultList []SubmitJobsJobResult

func (list *SubmitJobsJobResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsJobResult)
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

type SubmitJobsWaterMarkList []SubmitJobsWaterMark

func (list *SubmitJobsWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsWaterMark)
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

type SubmitJobsMergeList []SubmitJobsMerge

func (list *SubmitJobsMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsMerge)
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

type SubmitJobsOpeningList []SubmitJobsOpening

func (list *SubmitJobsOpeningList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsOpening)
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

type SubmitJobsTailSlateList []SubmitJobsTailSlate

func (list *SubmitJobsTailSlateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsTailSlate)
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

type SubmitJobsVideoStreamList []SubmitJobsVideoStream

func (list *SubmitJobsVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsVideoStream)
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

type SubmitJobsAudioStreamList []SubmitJobsAudioStream

func (list *SubmitJobsAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsAudioStream)
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

type SubmitJobsSubtitleStreamList []SubmitJobsSubtitleStream

func (list *SubmitJobsSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsSubtitleStream)
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

type SubmitJobsSubtitleList []SubmitJobsSubtitle

func (list *SubmitJobsSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsSubtitle)
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

type SubmitJobsExtSubtitleList []SubmitJobsExtSubtitle

func (list *SubmitJobsExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsExtSubtitle)
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
