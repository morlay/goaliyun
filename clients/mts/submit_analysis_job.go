package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SubmitAnalysisJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AnalysisConfig       string `position:"Query" name:"AnalysisConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             string `position:"Query" name:"Priority"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitAnalysisJobRequest) Invoke(client goaliyun.Client) (*SubmitAnalysisJobResponse, error) {
	resp := &SubmitAnalysisJobResponse{}
	err := client.Request("mts", "SubmitAnalysisJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAnalysisJobResponse struct {
	RequestId   goaliyun.String
	AnalysisJob SubmitAnalysisJobAnalysisJob
}

type SubmitAnalysisJobAnalysisJob struct {
	Id               goaliyun.String
	UserData         goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	Percent          goaliyun.Integer
	CreationTime     goaliyun.String
	PipelineId       goaliyun.String
	Priority         goaliyun.String
	TemplateList     SubmitAnalysisJobTemplateList
	InputFile        SubmitAnalysisJobInputFile
	AnalysisConfig   SubmitAnalysisJobAnalysisConfig
	MNSMessageResult SubmitAnalysisJobMNSMessageResult
}

type SubmitAnalysisJobTemplate struct {
	Id          goaliyun.String
	Name        goaliyun.String
	State       goaliyun.String
	Container   SubmitAnalysisJobContainer
	Video       SubmitAnalysisJobVideo
	Audio       SubmitAnalysisJobAudio
	TransConfig SubmitAnalysisJobTransConfig
	MuxConfig   SubmitAnalysisJobMuxConfig
}

type SubmitAnalysisJobContainer struct {
	Format goaliyun.String
}

type SubmitAnalysisJobVideo struct {
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
	BitrateBnd SubmitAnalysisJobBitrateBnd
}

type SubmitAnalysisJobBitrateBnd struct {
	Max goaliyun.String
	Min goaliyun.String
}

type SubmitAnalysisJobAudio struct {
	Codec      goaliyun.String
	Profile    goaliyun.String
	Samplerate goaliyun.String
	Bitrate    goaliyun.String
	Channels   goaliyun.String
	Qscale     goaliyun.String
}

type SubmitAnalysisJobTransConfig struct {
	TransMode goaliyun.String
}

type SubmitAnalysisJobMuxConfig struct {
	Segment SubmitAnalysisJobSegment
	Gif     SubmitAnalysisJobGif
}

type SubmitAnalysisJobSegment struct {
	Duration goaliyun.String
}

type SubmitAnalysisJobGif struct {
	Loop       goaliyun.String
	FinalDelay goaliyun.String
}

type SubmitAnalysisJobInputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type SubmitAnalysisJobAnalysisConfig struct {
	QualityControl    SubmitAnalysisJobQualityControl
	PropertiesControl SubmitAnalysisJobPropertiesControl
}

type SubmitAnalysisJobQualityControl struct {
	RateQuality     goaliyun.String
	MethodStreaming goaliyun.String
}

type SubmitAnalysisJobPropertiesControl struct {
	Deinterlace goaliyun.String
	Crop        SubmitAnalysisJobCrop
}

type SubmitAnalysisJobCrop struct {
	Mode   goaliyun.String
	Width  goaliyun.String
	Height goaliyun.String
	Top    goaliyun.String
	Left   goaliyun.String
}

type SubmitAnalysisJobMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type SubmitAnalysisJobTemplateList []SubmitAnalysisJobTemplate

func (list *SubmitAnalysisJobTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitAnalysisJobTemplate)
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
