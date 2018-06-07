package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SubmitEditingJobsRequest struct {
	OutputBucket         string `position:"Query" name:"OutputBucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	EditingJobOutputs    string `position:"Query" name:"EditingJobOutputs"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OutputLocation       string `position:"Query" name:"OutputLocation"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	EditingInputs        string `position:"Query" name:"EditingInputs"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitEditingJobsRequest) Invoke(client goaliyun.Client) (*SubmitEditingJobsResponse, error) {
	resp := &SubmitEditingJobsResponse{}
	err := client.Request("mts", "SubmitEditingJobs", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitEditingJobsResponse struct {
	RequestId     goaliyun.String
	JobResultList SubmitEditingJobsJobResultList
}

type SubmitEditingJobsJobResult struct {
	Success bool
	Code    goaliyun.String
	Message goaliyun.String
	Job     SubmitEditingJobsJob
}

type SubmitEditingJobsJob struct {
	JobId            goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	Percent          goaliyun.Integer
	PipelineId       goaliyun.String
	CreationTime     goaliyun.String
	FinishTime       goaliyun.String
	EditingInputs    SubmitEditingJobsEditingInputList
	EditingConfig    SubmitEditingJobsEditingConfig
	MNSMessageResult SubmitEditingJobsMNSMessageResult
}

type SubmitEditingJobsEditingInput struct {
	Id          goaliyun.String
	InputFile   SubmitEditingJobsInputFile
	InputConfig SubmitEditingJobsInputConfig
}

type SubmitEditingJobsInputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type SubmitEditingJobsInputConfig struct {
	DeinterlaceMethod goaliyun.String
	IsNormalSar       goaliyun.String
}

type SubmitEditingJobsEditingConfig struct {
	TemplateId             goaliyun.String
	UserData               goaliyun.String
	Rotate                 goaliyun.String
	VideoStreamMap         goaliyun.String
	AudioStreamMap         goaliyun.String
	DeWatermark            goaliyun.String
	Priority               goaliyun.String
	WaterMarkConfigUrl     goaliyun.String
	MergeConfigUrl         goaliyun.String
	WaterMarkList          SubmitEditingJobsWaterMarkList
	MergeList              SubmitEditingJobsMergeList
	DigiWaterMark          SubmitEditingJobsDigiWaterMark
	OutputFile             SubmitEditingJobsOutputFile
	M3U8NonStandardSupport SubmitEditingJobsM3U8NonStandardSupport
	Properties             SubmitEditingJobsProperties
	Clip                   SubmitEditingJobsClip
	SuperReso              SubmitEditingJobsSuperReso
	SubtitleConfig         SubmitEditingJobsSubtitleConfig
	TransConfig            SubmitEditingJobsTransConfig
	MuxConfig              SubmitEditingJobsMuxConfig
	Audio                  SubmitEditingJobsAudio
	Video                  SubmitEditingJobsVideo
	Container              SubmitEditingJobsContainer
	Encryption             SubmitEditingJobsEncryption
	Editing                SubmitEditingJobsEditing
}

type SubmitEditingJobsWaterMark struct {
	WaterMarkTemplateId goaliyun.String
	Width               goaliyun.String
	Height              goaliyun.String
	Dx                  goaliyun.String
	Dy                  goaliyun.String
	ReferPos            goaliyun.String
	Type                goaliyun.String
	InputFile1          SubmitEditingJobsInputFile1
}

type SubmitEditingJobsInputFile1 struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type SubmitEditingJobsMerge struct {
	MergeURL goaliyun.String
	Start    goaliyun.String
	Duration goaliyun.String
	RoleArn  goaliyun.String
}

type SubmitEditingJobsDigiWaterMark struct {
	Type       goaliyun.String
	Alpha      goaliyun.String
	InputFile2 SubmitEditingJobsInputFile2
}

type SubmitEditingJobsInputFile2 struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type SubmitEditingJobsOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type SubmitEditingJobsM3U8NonStandardSupport struct {
	TS SubmitEditingJobsTS
}

type SubmitEditingJobsTS struct {
	Md5Support  bool
	SizeSupport bool
}

type SubmitEditingJobsProperties struct {
	Width      goaliyun.String
	Height     goaliyun.String
	Bitrate    goaliyun.String
	Duration   goaliyun.String
	Fps        goaliyun.String
	FileSize   goaliyun.String
	FileFormat goaliyun.String
	Streams    SubmitEditingJobsStreams
	Format     SubmitEditingJobsFormat
}

type SubmitEditingJobsStreams struct {
	VideoStreamList    SubmitEditingJobsVideoStreamList
	AudioStreamList    SubmitEditingJobsAudioStreamList
	SubtitleStreamList SubmitEditingJobsSubtitleStreamList
}

type SubmitEditingJobsVideoStream struct {
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
	NetworkCost    SubmitEditingJobsNetworkCost
}

type SubmitEditingJobsNetworkCost struct {
	PreloadTime   goaliyun.String
	CostBandwidth goaliyun.String
	AvgBitrate    goaliyun.String
}

type SubmitEditingJobsAudioStream struct {
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

type SubmitEditingJobsSubtitleStream struct {
	Index goaliyun.String
	Lang  goaliyun.String
}

type SubmitEditingJobsFormat struct {
	NumStreams     goaliyun.String
	NumPrograms    goaliyun.String
	FormatName     goaliyun.String
	FormatLongName goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Size           goaliyun.String
	Bitrate        goaliyun.String
}

type SubmitEditingJobsClip struct {
	TimeSpan SubmitEditingJobsTimeSpan
}

type SubmitEditingJobsTimeSpan struct {
	Seek     goaliyun.String
	Duration goaliyun.String
}

type SubmitEditingJobsSuperReso struct {
	IsHalfSample goaliyun.String
}

type SubmitEditingJobsSubtitleConfig struct {
	SubtitleList    SubmitEditingJobsSubtitleList
	ExtSubtitleList SubmitEditingJobsExtSubtitleList
}

type SubmitEditingJobsSubtitle struct {
	Map goaliyun.String
}

type SubmitEditingJobsExtSubtitle struct {
	FontName goaliyun.String
	CharEnc  goaliyun.String
	Input    SubmitEditingJobsInput
}

type SubmitEditingJobsInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type SubmitEditingJobsTransConfig struct {
	TransMode               goaliyun.String
	IsCheckReso             goaliyun.String
	IsCheckResoFail         goaliyun.String
	IsCheckVideoBitrate     goaliyun.String
	IsCheckAudioBitrate     goaliyun.String
	AdjDarMethod            goaliyun.String
	IsCheckVideoBitrateFail goaliyun.String
	IsCheckAudioBitrateFail goaliyun.String
	Duration                goaliyun.String
}

type SubmitEditingJobsMuxConfig struct {
	Segment SubmitEditingJobsSegment
	Gif     SubmitEditingJobsGif
}

type SubmitEditingJobsSegment struct {
	Duration goaliyun.String
}

type SubmitEditingJobsGif struct {
	Loop            goaliyun.String
	FinalDelay      goaliyun.String
	IsCustomPalette goaliyun.String
	DitherMode      goaliyun.String
}

type SubmitEditingJobsAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
	Volume     SubmitEditingJobsVolume
}

type SubmitEditingJobsVolume struct {
	Level  goaliyun.String
	Method goaliyun.String
}

type SubmitEditingJobsVideo struct {
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
	BitrateBnd SubmitEditingJobsBitrateBnd
}

type SubmitEditingJobsBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type SubmitEditingJobsContainer struct {
	Format goaliyun.String
}

type SubmitEditingJobsEncryption struct {
	Type    goaliyun.String
	Id      goaliyun.String
	Key     goaliyun.String
	KeyUri  goaliyun.String
	KeyType goaliyun.String
	SkipCnt goaliyun.String
}

type SubmitEditingJobsEditing struct {
	ClipList SubmitEditingJobsClip3List
	Timeline SubmitEditingJobsTimeline
}

type SubmitEditingJobsClip3 struct {
	Id            goaliyun.String
	Type          goaliyun.String
	SourceType    goaliyun.String
	SourceID      goaliyun.String
	SourceStrmMap goaliyun.String
	In            goaliyun.String
	Out           goaliyun.String
	Effects       SubmitEditingJobsEffectList
}

type SubmitEditingJobsEffect struct {
	Effect       goaliyun.String
	EffectConfig goaliyun.String
}

type SubmitEditingJobsTimeline struct {
	TrackList      SubmitEditingJobsTrackList
	TimelineConfig SubmitEditingJobsTimelineConfig
}

type SubmitEditingJobsTrack struct {
	Id    goaliyun.String
	Type  goaliyun.String
	Order goaliyun.String
	Clips SubmitEditingJobsClip4List
}

type SubmitEditingJobsClip4 struct {
	ClipID      goaliyun.String
	In          goaliyun.String
	Out         goaliyun.String
	ClipsConfig SubmitEditingJobsClipsConfig
}

type SubmitEditingJobsClipsConfig struct {
	ClipsConfigVideo SubmitEditingJobsClipsConfigVideo
}

type SubmitEditingJobsClipsConfigVideo struct {
	L goaliyun.String
	T goaliyun.String
}

type SubmitEditingJobsTimelineConfig struct {
	TimelineConfigVideo SubmitEditingJobsTimelineConfigVideo
	TimelineConfigAudio SubmitEditingJobsTimelineConfigAudio
}

type SubmitEditingJobsTimelineConfigVideo struct {
	Width   goaliyun.String
	Height  goaliyun.String
	BgColor goaliyun.String
	Fps     goaliyun.String
}

type SubmitEditingJobsTimelineConfigAudio struct {
	Samplerate    goaliyun.String
	ChannelLayout goaliyun.String
	Channels      goaliyun.String
}

type SubmitEditingJobsMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type SubmitEditingJobsJobResultList []SubmitEditingJobsJobResult

func (list *SubmitEditingJobsJobResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsJobResult)
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

type SubmitEditingJobsEditingInputList []SubmitEditingJobsEditingInput

func (list *SubmitEditingJobsEditingInputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsEditingInput)
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

type SubmitEditingJobsWaterMarkList []SubmitEditingJobsWaterMark

func (list *SubmitEditingJobsWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsWaterMark)
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

type SubmitEditingJobsMergeList []SubmitEditingJobsMerge

func (list *SubmitEditingJobsMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsMerge)
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

type SubmitEditingJobsVideoStreamList []SubmitEditingJobsVideoStream

func (list *SubmitEditingJobsVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsVideoStream)
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

type SubmitEditingJobsAudioStreamList []SubmitEditingJobsAudioStream

func (list *SubmitEditingJobsAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsAudioStream)
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

type SubmitEditingJobsSubtitleStreamList []SubmitEditingJobsSubtitleStream

func (list *SubmitEditingJobsSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsSubtitleStream)
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

type SubmitEditingJobsSubtitleList []SubmitEditingJobsSubtitle

func (list *SubmitEditingJobsSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsSubtitle)
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

type SubmitEditingJobsExtSubtitleList []SubmitEditingJobsExtSubtitle

func (list *SubmitEditingJobsExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsExtSubtitle)
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

type SubmitEditingJobsClip3List []SubmitEditingJobsClip3

func (list *SubmitEditingJobsClip3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsClip3)
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

type SubmitEditingJobsEffectList []SubmitEditingJobsEffect

func (list *SubmitEditingJobsEffectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsEffect)
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

type SubmitEditingJobsTrackList []SubmitEditingJobsTrack

func (list *SubmitEditingJobsTrackList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsTrack)
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

type SubmitEditingJobsClip4List []SubmitEditingJobsClip4

func (list *SubmitEditingJobsClip4List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsClip4)
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
