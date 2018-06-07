package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryEditingJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryEditingJobListRequest) Invoke(client goaliyun.Client) (*QueryEditingJobListResponse, error) {
	resp := &QueryEditingJobListResponse{}
	err := client.Request("mts", "QueryEditingJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryEditingJobListResponse struct {
	RequestId      goaliyun.String
	JobList        QueryEditingJobListJobList
	NonExistJobIds QueryEditingJobListNonExistJobIdList
}

type QueryEditingJobListJob struct {
	JobId            goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	Percent          goaliyun.Integer
	PipelineId       goaliyun.String
	CreationTime     goaliyun.String
	FinishTime       goaliyun.String
	EditingInputs    QueryEditingJobListEditingInputList
	EditingConfig    QueryEditingJobListEditingConfig
	MNSMessageResult QueryEditingJobListMNSMessageResult
}

type QueryEditingJobListEditingInput struct {
	Id          goaliyun.String
	InputFile   QueryEditingJobListInputFile
	InputConfig QueryEditingJobListInputConfig
}

type QueryEditingJobListInputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryEditingJobListInputConfig struct {
	DeinterlaceMethod goaliyun.String
	IsNormalSar       goaliyun.String
}

type QueryEditingJobListEditingConfig struct {
	TemplateId             goaliyun.String
	UserData               goaliyun.String
	Rotate                 goaliyun.String
	VideoStreamMap         goaliyun.String
	AudioStreamMap         goaliyun.String
	DeWatermark            goaliyun.String
	Priority               goaliyun.String
	WaterMarkConfigUrl     goaliyun.String
	MergeConfigUrl         goaliyun.String
	WaterMarkList          QueryEditingJobListWaterMarkList
	MergeList              QueryEditingJobListMergeList
	DigiWaterMark          QueryEditingJobListDigiWaterMark
	OutputFile             QueryEditingJobListOutputFile
	M3U8NonStandardSupport QueryEditingJobListM3U8NonStandardSupport
	Properties             QueryEditingJobListProperties
	Clip                   QueryEditingJobListClip
	SuperReso              QueryEditingJobListSuperReso
	SubtitleConfig         QueryEditingJobListSubtitleConfig
	TransConfig            QueryEditingJobListTransConfig
	MuxConfig              QueryEditingJobListMuxConfig
	Audio                  QueryEditingJobListAudio
	Video                  QueryEditingJobListVideo
	Container              QueryEditingJobListContainer
	Encryption             QueryEditingJobListEncryption
	Editing                QueryEditingJobListEditing
}

type QueryEditingJobListWaterMark struct {
	WaterMarkTemplateId goaliyun.String
	Width               goaliyun.String
	Height              goaliyun.String
	Dx                  goaliyun.String
	Dy                  goaliyun.String
	ReferPos            goaliyun.String
	Type                goaliyun.String
	InputFile1          QueryEditingJobListInputFile1
}

type QueryEditingJobListInputFile1 struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryEditingJobListMerge struct {
	MergeURL goaliyun.String
	Start    goaliyun.String
	Duration goaliyun.String
	RoleArn  goaliyun.String
}

type QueryEditingJobListDigiWaterMark struct {
	Type       goaliyun.String
	Alpha      goaliyun.String
	InputFile2 QueryEditingJobListInputFile2
}

type QueryEditingJobListInputFile2 struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryEditingJobListOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type QueryEditingJobListM3U8NonStandardSupport struct {
	TS QueryEditingJobListTS
}

type QueryEditingJobListTS struct {
	Md5Support  bool
	SizeSupport bool
}

type QueryEditingJobListProperties struct {
	Width      goaliyun.String
	Height     goaliyun.String
	Bitrate    goaliyun.String
	Duration   goaliyun.String
	Fps        goaliyun.String
	FileSize   goaliyun.String
	FileFormat goaliyun.String
	Streams    QueryEditingJobListStreams
	Format     QueryEditingJobListFormat
}

type QueryEditingJobListStreams struct {
	VideoStreamList    QueryEditingJobListVideoStreamList
	AudioStreamList    QueryEditingJobListAudioStreamList
	SubtitleStreamList QueryEditingJobListSubtitleStreamList
}

type QueryEditingJobListVideoStream struct {
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
	NetworkCost    QueryEditingJobListNetworkCost
}

type QueryEditingJobListNetworkCost struct {
	PreloadTime   goaliyun.String
	CostBandwidth goaliyun.String
	AvgBitrate    goaliyun.String
}

type QueryEditingJobListAudioStream struct {
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

type QueryEditingJobListSubtitleStream struct {
	Index goaliyun.String
	Lang  goaliyun.String
}

type QueryEditingJobListFormat struct {
	NumStreams     goaliyun.String
	NumPrograms    goaliyun.String
	FormatName     goaliyun.String
	FormatLongName goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Size           goaliyun.String
	Bitrate        goaliyun.String
}

type QueryEditingJobListClip struct {
	TimeSpan QueryEditingJobListTimeSpan
}

type QueryEditingJobListTimeSpan struct {
	Seek     goaliyun.String
	Duration goaliyun.String
}

type QueryEditingJobListSuperReso struct {
	IsHalfSample goaliyun.String
}

type QueryEditingJobListSubtitleConfig struct {
	SubtitleList    QueryEditingJobListSubtitleList
	ExtSubtitleList QueryEditingJobListExtSubtitleList
}

type QueryEditingJobListSubtitle struct {
	Map goaliyun.String
}

type QueryEditingJobListExtSubtitle struct {
	FontName goaliyun.String
	CharEnc  goaliyun.String
	Input    QueryEditingJobListInput
}

type QueryEditingJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryEditingJobListTransConfig struct {
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

type QueryEditingJobListMuxConfig struct {
	Segment QueryEditingJobListSegment
	Gif     QueryEditingJobListGif
}

type QueryEditingJobListSegment struct {
	Duration goaliyun.String
}

type QueryEditingJobListGif struct {
	Loop            goaliyun.String
	FinalDelay      goaliyun.String
	IsCustomPalette goaliyun.String
	DitherMode      goaliyun.String
}

type QueryEditingJobListAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
	Volume     QueryEditingJobListVolume
}

type QueryEditingJobListVolume struct {
	Level  goaliyun.String
	Method goaliyun.String
}

type QueryEditingJobListVideo struct {
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
	BitrateBnd QueryEditingJobListBitrateBnd
}

type QueryEditingJobListBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type QueryEditingJobListContainer struct {
	Format goaliyun.String
}

type QueryEditingJobListEncryption struct {
	Type    goaliyun.String
	Id      goaliyun.String
	Key     goaliyun.String
	KeyUri  goaliyun.String
	KeyType goaliyun.String
	SkipCnt goaliyun.String
}

type QueryEditingJobListEditing struct {
	ClipList QueryEditingJobListClip3List
	Timeline QueryEditingJobListTimeline
}

type QueryEditingJobListClip3 struct {
	Id            goaliyun.String
	Type          goaliyun.String
	SourceType    goaliyun.String
	SourceID      goaliyun.String
	SourceStrmMap goaliyun.String
	In            goaliyun.String
	Out           goaliyun.String
	Effects       QueryEditingJobListEffectList
}

type QueryEditingJobListEffect struct {
	Effect       goaliyun.String
	EffectConfig goaliyun.String
}

type QueryEditingJobListTimeline struct {
	TrackList      QueryEditingJobListTrackList
	TimelineConfig QueryEditingJobListTimelineConfig
}

type QueryEditingJobListTrack struct {
	Id    goaliyun.String
	Type  goaliyun.String
	Order goaliyun.String
	Clips QueryEditingJobListClip4List
}

type QueryEditingJobListClip4 struct {
	ClipID      goaliyun.String
	In          goaliyun.String
	Out         goaliyun.String
	ClipsConfig QueryEditingJobListClipsConfig
}

type QueryEditingJobListClipsConfig struct {
	ClipsConfigVideo QueryEditingJobListClipsConfigVideo
}

type QueryEditingJobListClipsConfigVideo struct {
	L goaliyun.String
	T goaliyun.String
}

type QueryEditingJobListTimelineConfig struct {
	TimelineConfigVideo QueryEditingJobListTimelineConfigVideo
	TimelineConfigAudio QueryEditingJobListTimelineConfigAudio
}

type QueryEditingJobListTimelineConfigVideo struct {
	Width   goaliyun.String
	Height  goaliyun.String
	BgColor goaliyun.String
	Fps     goaliyun.String
}

type QueryEditingJobListTimelineConfigAudio struct {
	Samplerate    goaliyun.String
	ChannelLayout goaliyun.String
	Channels      goaliyun.String
}

type QueryEditingJobListMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type QueryEditingJobListJobList []QueryEditingJobListJob

func (list *QueryEditingJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListJob)
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

type QueryEditingJobListNonExistJobIdList []goaliyun.String

func (list *QueryEditingJobListNonExistJobIdList) UnmarshalJSON(data []byte) error {
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

type QueryEditingJobListEditingInputList []QueryEditingJobListEditingInput

func (list *QueryEditingJobListEditingInputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListEditingInput)
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

type QueryEditingJobListWaterMarkList []QueryEditingJobListWaterMark

func (list *QueryEditingJobListWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListWaterMark)
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

type QueryEditingJobListMergeList []QueryEditingJobListMerge

func (list *QueryEditingJobListMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListMerge)
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

type QueryEditingJobListVideoStreamList []QueryEditingJobListVideoStream

func (list *QueryEditingJobListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListVideoStream)
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

type QueryEditingJobListAudioStreamList []QueryEditingJobListAudioStream

func (list *QueryEditingJobListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListAudioStream)
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

type QueryEditingJobListSubtitleStreamList []QueryEditingJobListSubtitleStream

func (list *QueryEditingJobListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListSubtitleStream)
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

type QueryEditingJobListSubtitleList []QueryEditingJobListSubtitle

func (list *QueryEditingJobListSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListSubtitle)
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

type QueryEditingJobListExtSubtitleList []QueryEditingJobListExtSubtitle

func (list *QueryEditingJobListExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListExtSubtitle)
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

type QueryEditingJobListClip3List []QueryEditingJobListClip3

func (list *QueryEditingJobListClip3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListClip3)
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

type QueryEditingJobListEffectList []QueryEditingJobListEffect

func (list *QueryEditingJobListEffectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListEffect)
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

type QueryEditingJobListTrackList []QueryEditingJobListTrack

func (list *QueryEditingJobListTrackList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListTrack)
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

type QueryEditingJobListClip4List []QueryEditingJobListClip4

func (list *QueryEditingJobListClip4List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListClip4)
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
