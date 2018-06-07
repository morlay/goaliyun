package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryJobListRequest) Invoke(client goaliyun.Client) (*QueryJobListResponse, error) {
	resp := &QueryJobListResponse{}
	err := client.Request("mts", "QueryJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryJobListResponse struct {
	RequestId      goaliyun.String
	JobList        QueryJobListJobList
	NonExistJobIds QueryJobListNonExistJobIdList
}

type QueryJobListJob struct {
	JobId            goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	Percent          goaliyun.Integer
	PipelineId       goaliyun.String
	CreationTime     goaliyun.String
	FinishTime       goaliyun.String
	Input            QueryJobListInput
	Output           QueryJobListOutput
	MNSMessageResult QueryJobListMNSMessageResult
}

type QueryJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryJobListOutput struct {
	TemplateId             goaliyun.String
	UserData               goaliyun.String
	Rotate                 goaliyun.String
	VideoStreamMap         goaliyun.String
	AudioStreamMap         goaliyun.String
	DeWatermark            goaliyun.String
	Priority               goaliyun.String
	WaterMarkConfigUrl     goaliyun.String
	MergeConfigUrl         goaliyun.String
	WaterMarkList          QueryJobListWaterMarkList
	MergeList              QueryJobListMergeList
	OpeningList            QueryJobListOpeningList
	TailSlateList          QueryJobListTailSlateList
	OutputFile             QueryJobListOutputFile
	M3U8NonStandardSupport QueryJobListM3U8NonStandardSupport
	Properties             QueryJobListProperties
	Clip                   QueryJobListClip
	SuperReso              QueryJobListSuperReso
	SubtitleConfig         QueryJobListSubtitleConfig
	TransConfig            QueryJobListTransConfig
	MuxConfig              QueryJobListMuxConfig
	Audio                  QueryJobListAudio
	Video                  QueryJobListVideo
	Container              QueryJobListContainer
	Encryption             QueryJobListEncryption
}

type QueryJobListWaterMark struct {
	WaterMarkTemplateId goaliyun.String
	Width               goaliyun.String
	Height              goaliyun.String
	Dx                  goaliyun.String
	Dy                  goaliyun.String
	ReferPos            goaliyun.String
	Type                goaliyun.String
	InputFile           QueryJobListInputFile
}

type QueryJobListInputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryJobListMerge struct {
	MergeURL goaliyun.String
	Start    goaliyun.String
	Duration goaliyun.String
	RoleArn  goaliyun.String
}

type QueryJobListOpening struct {
	OpenUrl goaliyun.String
	Start   goaliyun.String
	Width   goaliyun.String
	Height  goaliyun.String
}

type QueryJobListTailSlate struct {
	TailUrl       goaliyun.String
	Start         goaliyun.String
	BlendDuration goaliyun.String
	Width         goaliyun.String
	Height        goaliyun.String
	IsMergeAudio  bool
	BgColor       goaliyun.String
}

type QueryJobListOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type QueryJobListM3U8NonStandardSupport struct {
	TS QueryJobListTS
}

type QueryJobListTS struct {
	Md5Support  bool
	SizeSupport bool
}

type QueryJobListProperties struct {
	Width       goaliyun.String
	Height      goaliyun.String
	Bitrate     goaliyun.String
	Duration    goaliyun.String
	Fps         goaliyun.String
	FileSize    goaliyun.String
	FileFormat  goaliyun.String
	SourceLogos QueryJobListSourceLogoList
	Streams     QueryJobListStreams
	Format      QueryJobListFormat
}

type QueryJobListSourceLogo struct {
	Source goaliyun.String
}

type QueryJobListStreams struct {
	VideoStreamList    QueryJobListVideoStreamList
	AudioStreamList    QueryJobListAudioStreamList
	SubtitleStreamList QueryJobListSubtitleStreamList
}

type QueryJobListVideoStream struct {
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
	NetworkCost    QueryJobListNetworkCost
}

type QueryJobListNetworkCost struct {
	PreloadTime   goaliyun.String
	CostBandwidth goaliyun.String
	AvgBitrate    goaliyun.String
}

type QueryJobListAudioStream struct {
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

type QueryJobListSubtitleStream struct {
	Index goaliyun.String
	Lang  goaliyun.String
}

type QueryJobListFormat struct {
	NumStreams     goaliyun.String
	NumPrograms    goaliyun.String
	FormatName     goaliyun.String
	FormatLongName goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Size           goaliyun.String
	Bitrate        goaliyun.String
}

type QueryJobListClip struct {
	TimeSpan QueryJobListTimeSpan
}

type QueryJobListTimeSpan struct {
	Seek     goaliyun.String
	Duration goaliyun.String
}

type QueryJobListSuperReso struct {
	IsHalfSample goaliyun.String
}

type QueryJobListSubtitleConfig struct {
	SubtitleList    QueryJobListSubtitleList
	ExtSubtitleList QueryJobListExtSubtitleList
}

type QueryJobListSubtitle struct {
	Map goaliyun.String
}

type QueryJobListExtSubtitle struct {
	FontName goaliyun.String
	CharEnc  goaliyun.String
	Input1   QueryJobListInput1
}

type QueryJobListInput1 struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryJobListTransConfig struct {
	TransMode               goaliyun.String
	IsCheckReso             goaliyun.String
	IsCheckResoFail         goaliyun.String
	IsCheckVideoBitrate     goaliyun.String
	IsCheckAudioBitrate     goaliyun.String
	AdjDarMethod            goaliyun.String
	IsCheckVideoBitrateFail goaliyun.String
	IsCheckAudioBitrateFail goaliyun.String
}

type QueryJobListMuxConfig struct {
	Segment QueryJobListSegment
	Gif     QueryJobListGif
}

type QueryJobListSegment struct {
	Duration goaliyun.String
}

type QueryJobListGif struct {
	Loop            goaliyun.String
	FinalDelay      goaliyun.String
	IsCustomPalette goaliyun.String
	DitherMode      goaliyun.String
}

type QueryJobListAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
	Volume     QueryJobListVolume
}

type QueryJobListVolume struct {
	Level  goaliyun.String
	Method goaliyun.String
}

type QueryJobListVideo struct {
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
	BitrateBnd QueryJobListBitrateBnd
}

type QueryJobListBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type QueryJobListContainer struct {
	Format goaliyun.String
}

type QueryJobListEncryption struct {
	Type    goaliyun.String
	Id      goaliyun.String
	Key     goaliyun.String
	KeyUri  goaliyun.String
	KeyType goaliyun.String
	SkipCnt goaliyun.String
}

type QueryJobListMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type QueryJobListJobList []QueryJobListJob

func (list *QueryJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListJob)
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

type QueryJobListNonExistJobIdList []goaliyun.String

func (list *QueryJobListNonExistJobIdList) UnmarshalJSON(data []byte) error {
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

type QueryJobListWaterMarkList []QueryJobListWaterMark

func (list *QueryJobListWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListWaterMark)
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

type QueryJobListMergeList []QueryJobListMerge

func (list *QueryJobListMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListMerge)
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

type QueryJobListOpeningList []QueryJobListOpening

func (list *QueryJobListOpeningList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListOpening)
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

type QueryJobListTailSlateList []QueryJobListTailSlate

func (list *QueryJobListTailSlateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListTailSlate)
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

type QueryJobListSourceLogoList []QueryJobListSourceLogo

func (list *QueryJobListSourceLogoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListSourceLogo)
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

type QueryJobListVideoStreamList []QueryJobListVideoStream

func (list *QueryJobListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListVideoStream)
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

type QueryJobListAudioStreamList []QueryJobListAudioStream

func (list *QueryJobListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListAudioStream)
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

type QueryJobListSubtitleStreamList []QueryJobListSubtitleStream

func (list *QueryJobListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListSubtitleStream)
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

type QueryJobListSubtitleList []QueryJobListSubtitle

func (list *QueryJobListSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListSubtitle)
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

type QueryJobListExtSubtitleList []QueryJobListExtSubtitle

func (list *QueryJobListExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListExtSubtitle)
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
