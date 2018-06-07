package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryAnalysisJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AnalysisJobIds       string `position:"Query" name:"AnalysisJobIds"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryAnalysisJobListRequest) Invoke(client goaliyun.Client) (*QueryAnalysisJobListResponse, error) {
	resp := &QueryAnalysisJobListResponse{}
	err := client.Request("mts", "QueryAnalysisJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryAnalysisJobListResponse struct {
	RequestId              goaliyun.String
	AnalysisJobList        QueryAnalysisJobListAnalysisJobList
	NonExistAnalysisJobIds QueryAnalysisJobListNonExistAnalysisJobIdList
}

type QueryAnalysisJobListAnalysisJob struct {
	Id               goaliyun.String
	UserData         goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	Percent          goaliyun.Integer
	CreationTime     goaliyun.String
	PipelineId       goaliyun.String
	Priority         goaliyun.String
	TemplateList     QueryAnalysisJobListTemplateList
	InputFile        QueryAnalysisJobListInputFile
	AnalysisConfig   QueryAnalysisJobListAnalysisConfig
	MNSMessageResult QueryAnalysisJobListMNSMessageResult
}

type QueryAnalysisJobListTemplate struct {
	Id          goaliyun.String
	Name        goaliyun.String
	State       goaliyun.String
	Container   QueryAnalysisJobListContainer
	Video       QueryAnalysisJobListVideo
	Audio       QueryAnalysisJobListAudio
	TransConfig QueryAnalysisJobListTransConfig
	MuxConfig   QueryAnalysisJobListMuxConfig
}

type QueryAnalysisJobListContainer struct {
	Format goaliyun.String
}

type QueryAnalysisJobListVideo struct {
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
	BitrateBnd QueryAnalysisJobListBitrateBnd
}

type QueryAnalysisJobListBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type QueryAnalysisJobListAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
}

type QueryAnalysisJobListTransConfig struct {
	TransMode goaliyun.String
}

type QueryAnalysisJobListMuxConfig struct {
	Segment QueryAnalysisJobListSegment
	Gif     QueryAnalysisJobListGif
}

type QueryAnalysisJobListSegment struct {
	Duration goaliyun.String
}

type QueryAnalysisJobListGif struct {
	Loop       goaliyun.String
	FinalDelay goaliyun.String
}

type QueryAnalysisJobListInputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryAnalysisJobListAnalysisConfig struct {
	QualityControl    QueryAnalysisJobListQualityControl
	PropertiesControl QueryAnalysisJobListPropertiesControl
}

type QueryAnalysisJobListQualityControl struct {
	RateQuality     goaliyun.String
	MethodStreaming goaliyun.String
}

type QueryAnalysisJobListPropertiesControl struct {
	Deinterlace goaliyun.String
	Crop        QueryAnalysisJobListCrop
}

type QueryAnalysisJobListCrop struct {
	Mode   goaliyun.String
	Width  goaliyun.String
	Height goaliyun.String
	Top    goaliyun.String
	Left   goaliyun.String
}

type QueryAnalysisJobListMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type QueryAnalysisJobListAnalysisJobList []QueryAnalysisJobListAnalysisJob

func (list *QueryAnalysisJobListAnalysisJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnalysisJobListAnalysisJob)
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

type QueryAnalysisJobListNonExistAnalysisJobIdList []goaliyun.String

func (list *QueryAnalysisJobListNonExistAnalysisJobIdList) UnmarshalJSON(data []byte) error {
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

type QueryAnalysisJobListTemplateList []QueryAnalysisJobListTemplate

func (list *QueryAnalysisJobListTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnalysisJobListTemplate)
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
